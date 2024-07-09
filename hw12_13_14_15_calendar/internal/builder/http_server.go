package builder

import (
	"context"

	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/app"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/server/http"
)

func (b *Builder) CreateHTTPServer(events app.EventsUseCaseInterface) *http.Server {
	server := http.CreateHTTPServer(events, b.config.HTTPAddr())

	b.shutdown.add(
		func(ctx context.Context) error {
			server.Stop(ctx)

			return nil
		},
	)

	return server
}
