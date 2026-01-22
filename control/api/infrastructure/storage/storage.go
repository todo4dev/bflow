// infrastructure/storage/storage.go
package storage

import (
	"fmt"
	"slices"
	"src/infrastructure/storage/aws_s3"
	"src/port/storage"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
	"github.com/leandroluk/gox/util"
)

func Provide() {
	provider := env.Get("STORAGE_PROVIDER", "aws_s3")
	switch provider {
	case "aws_s3":
		config := aws_s3.Config{
			Region:         env.Get("STORAGE_REGION", "us-east-1"),
			Bucket:         env.Get("STORAGE_BUCKET", ""),
			AccessKey:      env.Get("STORAGE_ACCESS_KEY", ""),
			SecretKey:      env.Get("STORAGE_SECRET_KEY", ""),
			Endpoint:       util.Ptr(env.Get("STORAGE_ENDPOINT", "")),
			ForcePathStyle: slices.Contains([]string{"true", "1"}, env.Get("STORAGE_FORCE_PATH_STYLE", "false")),
		}
		if err := config.Validate(); err != nil {
			panic(fmt.Errorf("storage config validation failed: %w", err))
		}

		instance, err := aws_s3.NewClient(&config)
		if err != nil {
			panic(fmt.Errorf("failed to create storage client: %w", err))
		}

		di.SingletonAs[storage.Client](func() storage.Client { return instance })
	default:
		panic(fmt.Errorf("invalid 'STORAGE_PROVIDER': %s", provider))
	}
}

func getEndpointPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
