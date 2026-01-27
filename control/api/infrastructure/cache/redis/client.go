// infrastructure/cache/redis/client.go
package redis

import (
	"context"
	"src/port/cache"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
	ttl    time.Duration
}

var _ cache.Client = (*Client)(nil)

func New(config *Config) (*Client, error) {
	opts, err := redis.ParseURL(config.URL)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: redis.NewClient(opts),
		ttl:    config.TTL,
	}, nil
}

func (c *Client) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	if ttl == 0 {
		ttl = c.ttl
	}
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *Client) findKey(ctx context.Context, match string) (string, error) {
	// Use Scan to find a matching key safely
	iter := c.client.Scan(ctx, 0, match, 1).Iterator()
	if iter.Next(ctx) {
		return iter.Val(), nil
	}
	if err := iter.Err(); err != nil {
		return "", err
	}
	return "", redis.Nil
}

func (c *Client) Get(ctx context.Context, match string) (string, string, error) {
	key, err := c.findKey(ctx, match)
	if err != nil {
		return "", "", err
	}
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", "", err
	}
	return key, val, nil
}

func (c *Client) Delete(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

func (c *Client) Exists(ctx context.Context, keys ...string) (int64, error) {
	return c.client.Exists(ctx, keys...).Result()
}

func (c *Client) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return c.client.Expire(ctx, key, ttl).Err()
}

func (c *Client) Increment(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

func (c *Client) Decrement(ctx context.Context, key string) (int64, error) {
	return c.client.Decr(ctx, key).Result()
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}
