package repository

import (
	"context"

	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type EventRepositoryInterface interface {
	Get(ctx context.Context, eventID int64) (*entity.Event, error)
	GetByOwnerID(ctx context.Context, ownerID int64) ([]*entity.Event, error)
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, eventID int64) error
}
