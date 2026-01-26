// application/health/health.go
package health

import (
	"context"
	"time"

	"src/port/broker"
	"src/port/cache"
	"src/port/database"
	"src/port/storage"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
)

var startTime time.Time

type pinger interface {
	Ping(ctx context.Context) error
}

type Result struct {
	Uptime string `json:"uptime"`
}

type Handler struct {
	services []pinger
}

func New(
	databaseClient database.Client,
	cacheClient cache.Client,
	brokerClient broker.Client,
	storageClient storage.Client,
) *Handler {
	return &Handler{
		services: []pinger{
			brokerClient,
			cacheClient,
			databaseClient,
			storageClient,
		},
	}
}

func (u *Handler) Handle() (*Result, error) {
	for _, service := range u.services {
		if err := service.Ping(context.Background()); err != nil {
			return nil, err
		}
	}

	return &Result{Uptime: time.Since(startTime).String()}, nil
}

func init() {
	startTime = time.Now()

	result := Result{
		Uptime: time.Since(startTime).String()}
	meta.Describe(&result, meta.Description("Result of health check"),
		meta.Field(&result.Uptime, meta.Description("Uptime")),
		meta.Example(result))
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
