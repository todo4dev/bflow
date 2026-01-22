// application/system/healthcheck/healthcheck.go
package healthcheck

import (
	"context"
	"errors"
	"strings"
	"time"

	"src/port/broker"
	"src/port/cache"
	"src/port/database"
	"src/port/storage"

	"github.com/leandroluk/gox/meta"
)

var startTime time.Time

type pinger interface {
	Ping(ctx context.Context) error
}

type Result struct {
	Uptime   string           `json:"uptime"`
	Status   bool             `json:"status"`
	Services map[string]error `json:"services"`
}

type Handler struct {
	services map[string]pinger
}

func New(
	database database.Client,
	cache cache.Client,
	broker broker.Client[string],
	storage storage.Client,
) *Handler {
	return &Handler{
		services: map[string]pinger{
			"database": database,
			"cache":    cache,
			"broker":   broker,
			"storage":  storage,
		},
	}
}

func (u *Handler) Handle(ctx context.Context) (*Result, error) {
	result := Result{
		Uptime:   time.Since(startTime).String(),
		Status:   true,
		Services: map[string]error{},
	}

	for name, service := range u.services {
		result.Services[name] = service.Ping(ctx)
		if result.Services[name] != nil {
			result.Status = false
		}
	}

	if !result.Status {
		var message strings.Builder
		message.WriteString("[healthCheck failed] ")

		for _, err := range result.Services {
			if err != nil {
				message.WriteString(err.Error() + "\n")
			}
		}
		return nil, errors.New(message.String())
	}
	return &result, nil
}

func init() {
	startTime = time.Now()

	result := Result{
		Uptime: time.Since(startTime).String(),
		Status: true,
		Services: map[string]error{
			"database": nil,
			"cache":    errors.New("failed to connect"),
			"broker":   nil,
			"storage":  nil}}
	meta.Describe(&result, meta.Description("Result of health check"),
		meta.Field(&result.Uptime, meta.Description("Uptime of the application")),
		meta.Field(&result.Status, meta.Description("Status of the application")),
		meta.Field(&result.Services, meta.Description("Services of the application")),
		meta.Example(result))
}
