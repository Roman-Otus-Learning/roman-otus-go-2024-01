package builder

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // for pgx driver
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (b *Builder) CreatePgSQLConnection(ctx context.Context) (*sqlx.DB, error) {
	if b.connection != nil {
		return b.connection, nil
	}

	var err error

	b.connection, err = sqlx.Open("pgx", b.config.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("open pgx driver: %w", err)
	}

	err = b.connection.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping connection: %w", err)
	}

	log.Info().Msg("sql server started")
	b.shutdown.add(
		func(ctx context.Context) error {
			b.connection.Close() //nolint:golint
			log.Info().Msg("sql server stopped")

			return nil
		},
	)

	return b.connection, nil
}
