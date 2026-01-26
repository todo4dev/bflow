// infrastructure/storage/s3/s3.go
package s3

import (
	"fmt"
	"slices"
	"src/port/storage"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
	"github.com/leandroluk/gox/util"
)

func Provide() {
	config := Config{
		Region:         env.Get("API_STORAGE_S3_REGION", "us-east-1"),
		Bucket:         env.Get("API_STORAGE_S3_BUCKET", ""),
		AccessKey:      env.Get("API_STORAGE_S3_ACCESS_KEY", ""),
		SecretKey:      env.Get("API_STORAGE_S3_SECRET_KEY", ""),
		Endpoint:       util.Ptr(env.Get("API_STORAGE_S3_ENDPOINT", "")),
		ForcePathStyle: slices.Contains([]string{"true", "1"}, env.Get("API_STORAGE_S3_FORCE_PATH_STYLE", "false")),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("storage config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create storage client: %w", err))
	}

	di.SingletonAs[storage.Client](func() storage.Client { return instance })
}
