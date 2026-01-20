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

	"github.com/leandroluk/gox/di"
)

var startTime time.Time

func init() { startTime = time.Now() }

type Result struct {
	Uptime   time.Duration    `json:"uptime"`
	Status   bool             `json:"status"`
	Services map[string]error `json:"services"`
}

func UseCase(ctx context.Context) (*Result, error) {
	result := Result{
		Uptime:   time.Since(startTime),
		Status:   true,
		Services: map[string]error{},
	}
	clients := map[string]interface {
		Ping(ctx context.Context) error
	}{
		"database": di.Resolve[database.Client](),
		"cache":    di.Resolve[cache.Client](),
		"broker":   di.Resolve[broker.Client](),
		"storage":  di.Resolve[storage.Client](),
	}

	for name, client := range clients {
		result.Services[name] = client.Ping(ctx)
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
