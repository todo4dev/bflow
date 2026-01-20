// infrastructure/storage/aws_s3/config.go
package aws_s3

import v "github.com/leandroluk/gox/validate"

type Config struct {
	Region         string
	Bucket         string
	AccessKey      string
	SecretKey      string
	Endpoint       *string // Optional, required for MinIO
	ForcePathStyle bool    // Set true for MinIO
}

var ConfigSchema = v.Object(func(t *Config, s *v.ObjectSchema[Config]) {
	s.Field(&t.Region).Text().Default("us-east-1")
	s.Field(&t.Bucket).Text().Required()
	s.Field(&t.AccessKey).Text().Required()
	s.Field(&t.SecretKey).Text().Required()
	s.Field(&t.Endpoint).Text()
	s.Field(&t.ForcePathStyle).Boolean().Default(false)
})
