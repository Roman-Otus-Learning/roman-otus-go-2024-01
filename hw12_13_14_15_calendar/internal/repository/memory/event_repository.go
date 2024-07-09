package memory

import (
	"context"
	"sync"

	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/repository"
	"github.com/pkg/errors"
	"golang.org/x/exp/rand"
)

var ErrEventAlreadyExists = errors.New("event already exists")

type EventMemoryRepository struct {
	mu     sync.RWMutex
	events map[int64]*entity.Event
}

func CreateEventMemoryRepository() *EventMemoryRepository {
	return &EventMemoryRepository{
		events: make(map[int64]*entity.Event),
	}
}

func (r *EventMemoryRepository) Get(_ context.Context, eventID int64) (*entity.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	event, isExists := r.events[eventID]
	if !isExists {
		return nil, repository.ErrNotFound
	}

	return event, nil
}

func (r *EventMemoryRepository) GetByOwnerID(_ context.Context, ownerID int64) ([]*entity.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	events := make([]*entity.Event, 0)
	for _, event := range r.events {
		if event.OwnerID == ownerID {
			events = append(events, event)
		}
	}

	return events, nil
}

func (r *EventMemoryRepository) Create(ctx context.Context, event *entity.Event) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	eventID := int64(rand.Int())

	existingEvent, err := r.Get(ctx, eventID)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return err
		}
	}

	if existingEvent != nil {
		return ErrEventAlreadyExists
	}

	event.ID = eventID
	r.events[event.ID] = event

	return nil
}

func (r *EventMemoryRepository) Update(ctx context.Context, event *entity.Event) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, err := r.Get(ctx, event.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return err
		}
	}

	r.events[event.ID] = event

	return nil
}

func (r *EventMemoryRepository) Delete(_ context.Context, eventID int64) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	delete(r.events, eventID)

	return nil
}
