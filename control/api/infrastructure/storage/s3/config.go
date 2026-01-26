// infrastructure/storage/s3/config.go
package s3

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Region).Text().Default("us-east-1")
	s.Field(&t.Bucket).Text().Required()
	s.Field(&t.AccessKey).Text().Required()
	s.Field(&t.SecretKey).Text().Required()
	s.Field(&t.Endpoint).Text()
	s.Field(&t.ForcePathStyle).Boolean().Default(false)
})

type Config struct {
	Region         string
	Bucket         string
	AccessKey      string
	SecretKey      string
	Endpoint       *string // Optional, required for MinIO
	ForcePathStyle bool    // Set true for MinIO
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
