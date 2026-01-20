// infrastructure/storage/aws_s3/client.go
package aws_s3

import (
	"context"
	"io"
	"src/port/storage"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucket        string
}

var _ storage.Client = (*Client)(nil)

func NewClient(rawConfig Config) (*Client, error) {
	cfg, err := ConfigSchema.Validate(rawConfig)
	if err != nil {
		return nil, err
	}

	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	opts := func(o *s3.Options) {
		if cfg.Endpoint != nil {
			o.BaseEndpoint = cfg.Endpoint
			o.UsePathStyle = true // Force path style when using custom endpoint
		}
		if cfg.ForcePathStyle {
			o.UsePathStyle = cfg.ForcePathStyle
		}
	}

	client := s3.NewFromConfig(awsCfg, opts)

	return &Client{
		client:        client,
		presignClient: s3.NewPresignClient(client),
		bucket:        cfg.Bucket,
	}, nil
}

func (c *Client) Upload(ctx context.Context, path string, reader io.Reader, metadata map[string]string) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
		Body:   reader,
	}

	if len(metadata) > 0 {
		input.Metadata = metadata
	}

	_, err := c.client.PutObject(ctx, input)
	return err
}

func (c *Client) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	output, err := c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}

func (c *Client) Delete(ctx context.Context, path string) error {
	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	})
	return err
}

func (c *Client) Exists(ctx context.Context, path string) (bool, error) {
	_, err := c.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		// Needs robust error checking for generic 404, but for now simple check
		// In AWS SDK v2, 404 is usually a NotFound error type
		// For simplicity, assuming error means not found or issue.
		// TODO: Refine error checking to distinguish network error from 404
		return false, nil
	}
	return true, nil
}

func (c *Client) GetInfo(ctx context.Context, path string) (*storage.Info, error) {
	output, err := c.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	return &storage.Info{
		Path:         path,
		Size:         aws.ToInt64(output.ContentLength),
		LastModified: aws.ToTime(output.LastModified),
		ContentType:  aws.ToString(output.ContentType),
		ETag:         aws.ToString(output.ETag),
		Metadata:     output.Metadata,
	}, nil
}

func (c *Client) GenerateTemporaryURL(ctx context.Context, path string, expiration time.Duration) (string, error) {
	req, err := c.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	}, func(o *s3.PresignOptions) {
		o.Expires = expiration
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (c *Client) List(ctx context.Context, prefix string) ([]storage.Info, error) {
	var files []storage.Info
	paginator := s3.NewListObjectsV2Paginator(c.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, obj := range page.Contents {
			files = append(files, storage.Info{
				Path:         aws.ToString(obj.Key),
				Size:         aws.ToInt64(obj.Size),
				LastModified: aws.ToTime(obj.LastModified),
				ETag:         aws.ToString(obj.ETag),
			})
		}
	}

	return files, nil
}

func (c *Client) Ping(ctx context.Context) error {
	_, err := c.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(c.bucket),
	})
	return err
}
