// infrastructure/database/jackc_pgx/client.go
package jackc_pgx

import (
	"context"
	sqldb "src/port/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	pool *pgxpool.Pool
	tx   pgx.Tx
}

var _ sqldb.Client = (*Client)(nil)

func NewClient(config *Config) (*Client, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, config.DSN)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return &Client{pool: pool}, nil
}

// conn returns the executor (Pool or Tx)
func (c *Client) conn() interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
} {
	if c.tx != nil {
		return c.tx
	}
	return c.pool
}

func (c *Client) Query(ctx context.Context, query string, args ...any) (sqldb.Rows, error) {
	rows, err := c.conn().Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &Rows{rows: rows}, nil
}

func (c *Client) QueryRow(ctx context.Context, query string, args ...any) sqldb.Row {
	return c.conn().QueryRow(ctx, query, args...)
}

func (c *Client) Exec(ctx context.Context, query string, args ...any) (sqldb.Result, error) {
	tag, err := c.conn().Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &Result{tag: tag}, nil
}

func (c *Client) Transaction(ctx context.Context, fn func(tx sqldb.Client) error) error {
	if c.tx != nil {
		// Nested transaction? For now, just reuse existing tx or error?
		// pgx supports Savepoints for nested tx, but let's keep it simple: assume no deep nesting or just execute in current tx.
		return fn(c)
	}

	return pgx.BeginFunc(ctx, c.pool, func(tx pgx.Tx) error {
		return fn(&Client{tx: tx})
	})
}

func (c *Client) Close() error {
	if c.pool != nil {
		c.pool.Close()
	}
	return nil
}

func (c *Client) Ping(ctx context.Context) error {
	if c.pool != nil {
		return c.pool.Ping(ctx)
	}
	// If in tx, assuming alive?
	if c.tx != nil {
		return c.tx.Conn().Ping(ctx)
	}
	return nil
}
