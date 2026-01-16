// infrastructure/cache/redis/client.go
package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Client struct {
	raw *goredis.Client
}

func NewClient(config *Config) (*Client, error) {
	if _, err := ConfigSchema.Validate(&config); err != nil {
		return nil, err
	}

	opt, err := goredis.ParseURL(config.URL)
	if err != nil {
		return nil, err
	}

	return &Client{raw: goredis.NewClient(opt)}, nil
}

func (c *Client) Raw() *goredis.Client {
	return c.raw
}

func (c *Client) Ping(ctx context.Context) error {
	return c.raw.Ping(ctx).Err()
}

func (c *Client) Close() error {
	return c.raw.Close()
}

func (c *Client) WaitUntilReady(ctx context.Context, retryInterval time.Duration) error {
	ticker := time.NewTicker(retryInterval)
	defer ticker.Stop()

	for {
		if err := c.Ping(ctx); err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
