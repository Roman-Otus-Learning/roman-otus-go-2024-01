package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/repository"
	"github.com/jmoiron/sqlx"
)

type EventSQLRepositoryInterface interface {
	repository.EventRepositoryInterface
}

type EventRepository struct {
	db *sqlx.DB
}

func CreateEventRepository(db *sqlx.DB) EventSQLRepositoryInterface {
	return &EventRepository{db: db}
}

func (r *EventRepository) Get(ctx context.Context, eventID int64) (*entity.Event, error) {
	event := entity.Event{}
	query := "SELECT * FROM event where id = $1"

	err := r.db.GetContext(ctx, &event, query, eventID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, repository.ErrNotFound
		}

		return nil, fmt.Errorf("get event: %w", err)
	}
	return &event, nil
}

func (r *EventRepository) GetByOwnerID(ctx context.Context, ownerID int64) ([]*entity.Event, error) {
	query := "SELECT * FROM event where owner_id = $1"
	rows, err := r.db.QueryContext(ctx, query, ownerID)
	if err != nil {
		return nil, fmt.Errorf("get events by owner id: %w", err)
	}
	defer rows.Close() //nolint:golint

	var events []*entity.Event

	for rows.Next() {
		event := &entity.Event{}

		if err := rows.Scan(
			&event.ID,
			&event.OwnerID,
			&event.Title,
			&event.Description,
			&event.TimeStart,
			&event.TimeEnd,
			&event.NotificationTime,
		); err != nil {
			return nil, fmt.Errorf("scan events by owner id: %w", err)
		}

		events = append(events, event)
	}
	return events, nil
}

func (r *EventRepository) Create(ctx context.Context, event *entity.Event) error {
	query := `
		INSERT INTO 
			event (owner_id, title, description, time_start, time_end, notification_time)
		VALUES 
			(:owner_id, :title, :description, :time_start, :time_end, :notification_time)
		;
	`

	_, err := r.db.NamedExecContext(
		ctx,
		query,
		map[string]interface{}{
			"owner_id":          event.OwnerID,
			"title":             event.Title,
			"description":       event.Description,
			"time_start":        event.TimeStart,
			"time_end":          event.TimeEnd,
			"notification_time": event.NotificationTime,
		},
	)
	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}

	return nil
}

func (r *EventRepository) Update(ctx context.Context, event *entity.Event) error {
	query := `
		UPDATE 
			event
		SET 
			owner_id=:owner_id,
			title=:title,
			description=:description,
			time_start=:time_start,
			time_end=:time_end,
			notification_time=:notification_time
		WHERE
			id=:id
		;
	`
	_, err := r.db.NamedExecContext(
		ctx,
		query,
		map[string]interface{}{
			"id":                event.ID,
			"owner_id":          event.OwnerID,
			"title":             event.Title,
			"description":       event.Description,
			"time_start":        event.TimeStart,
			"time_end":          event.TimeEnd,
			"notification_time": event.NotificationTime,
		},
	)
	if err != nil {
		return fmt.Errorf("update event: %w", err)
	}

	return nil
}

func (r *EventRepository) Delete(ctx context.Context, eventID int64) error {
	query := "DELETE FROM event where id = $1;"
	_, err := r.db.ExecContext(ctx, query, eventID)
	if err != nil {
		return fmt.Errorf("delete event: %w", err)
	}
	return nil
}
