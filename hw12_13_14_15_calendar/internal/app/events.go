package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/repository"
)

type EventsUseCaseInterface interface {
	Get(ctx context.Context, eventID int64) (*entity.Event, error)
	GetByOwnerID(ctx context.Context, ownerID int64) ([]*entity.Event, error)
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type Events struct {
	repository repository.EventRepositoryInterface
}

func NewEventUseCase(repository repository.EventRepositoryInterface) EventsUseCaseInterface {
	return &Events{
		repository: repository,
	}
}

func (e *Events) Get(ctx context.Context, eventID int64) (*entity.Event, error) {
	event, err := e.repository.Get(ctx, eventID)
	if err != nil {
		return nil, fmt.Errorf("get event: %w", err)
	}

	return event, nil
}

func (e *Events) GetByOwnerID(ctx context.Context, ownerID int64) ([]*entity.Event, error) {
	events, err := e.repository.GetByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("get by owner id: %w", err)
	}

	return events, nil
}

func (e *Events) Create(ctx context.Context) error {
	timeNow := time.Now()
	event := &entity.Event{
		OwnerID:          1,
		Title:            "Тестовый заголовок",
		Description:      "Тестовое описание",
		TimeStart:        timeNow,
		TimeEnd:          timeNow,
		NotificationTime: sql.NullTime{},
	}

	err := e.repository.Create(ctx, event)
	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}

	return nil
}

func (e *Events) Update(ctx context.Context) error {
	events, err := e.GetByOwnerID(ctx, 1)
	if err != nil {
		return fmt.Errorf("get events: %w", err)
	}
	event := events[0]

	event.Title = "Тестовый заголовок измененный"
	event.Description = "Тестовое описание измененное"

	err = e.repository.Update(ctx, event)
	if err != nil {
		return fmt.Errorf("update event: %w", err)
	}

	return nil
}

func (e *Events) Delete(ctx context.Context) error {
	events, err := e.GetByOwnerID(ctx, 1)
	if err != nil {
		return fmt.Errorf("get events: %w", err)
	}

	if len(events) == 0 {
		return fmt.Errorf("not found event for delete")
	}

	event := events[len(events)-1]

	err = e.repository.Delete(ctx, event.ID)
	if err != nil {
		return fmt.Errorf("delete event: %w", err)
	}

	return nil
}
