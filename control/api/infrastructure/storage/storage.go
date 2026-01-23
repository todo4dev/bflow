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
	provider := env.Get("API_STORAGE_PROVIDER", "aws_s3")
	switch provider {
	case "aws_s3":
		config := aws_s3.Config{
			Region:         env.Get("API_STORAGE_REGION", "us-east-1"),
			Bucket:         env.Get("API_STORAGE_BUCKET", ""),
			AccessKey:      env.Get("API_STORAGE_ACCESS_KEY", ""),
			SecretKey:      env.Get("API_STORAGE_SECRET_KEY", ""),
			Endpoint:       util.Ptr(env.Get("API_STORAGE_ENDPOINT", "")),
			ForcePathStyle: slices.Contains([]string{"true", "1"}, env.Get("API_STORAGE_FORCE_PATH_STYLE", "false")),
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
		panic(fmt.Errorf("invalid 'API_STORAGE_PROVIDER': %s", provider))
	}
}
