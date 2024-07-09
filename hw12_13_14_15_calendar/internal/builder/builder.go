package builder

import (
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/app"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/config"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/repository"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/repository/database"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/repository/memory"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type Builder struct {
	config     *config.Config
	connection *sqlx.DB
	shutdown   shutdown
}

func CreateBuilder(config *config.Config) *Builder {
	return &Builder{config: config}
}

func (b *Builder) EventRepository(ctx context.Context) (repository.EventRepositoryInterface, error) {
	if b.config.Database.InMemory {
		log.Info().Msg("creating in-memory events repository")

		return memory.CreateEventMemoryRepository(), nil
	}

	pgSQLConnection, err := b.CreatePgSQLConnection(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get pgsql connection")
	}

	return database.CreateEventRepository(pgSQLConnection), nil
}

func (b *Builder) EventsUseCases(repository repository.EventRepositoryInterface) app.EventsUseCaseInterface {
	return app.NewEventUseCase(repository)
}
