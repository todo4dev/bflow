// infrastructure/storage/storage.go
package storage

import (
	"fmt"
	"slices"
	"src/infrastructure/storage/aws_s3"
	"src/port/storage"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	di.SingletonAs[storage.Client](func() storage.Client {
		provider := env.Get("STORAGE_PROVIDER", "aws_s3")
		switch provider {
		case "aws_s3":
			endpoint := env.Get("STORAGE_ENDPOINT", "")
			var endpointPtr *string
			if endpoint != "" {
				endpointPtr = &endpoint
			}

			config := aws_s3.Config{
				Region:         env.Get("STORAGE_REGION", "us-east-1"),
				Bucket:         env.Get("STORAGE_BUCKET", ""),
				AccessKey:      env.Get("STORAGE_ACCESS_KEY", ""),
				SecretKey:      env.Get("STORAGE_SECRET_KEY", ""),
				Endpoint:       endpointPtr,
				ForcePathStyle: slices.Contains([]string{"true", "1"}, env.Get("STORAGE_FORCE_PATH_STYLE", "false")),
			}
			if _, err := aws_s3.ConfigSchema.Validate(&config); err != nil {
				panic(fmt.Errorf("storage config validation failed: %w", err))
			}
			instance, err := aws_s3.NewClient(config)
			if err != nil {
				panic(fmt.Errorf("failed to create storage client: %w", err))
			}
			return instance
		default:
			panic(fmt.Errorf("invalid 'STORAGE_PROVIDER': %s", provider))
		}
	})
}
