package field

import (
	"context"
	"sync"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
)

var GlobFields = map[domain.FieldID]*domain.ImageID{
	1: nil,
	2: nil,
}

type Repository struct {
	mu     sync.RWMutex
	Fields map[domain.FieldID]*domain.ImageID
}

func NewRepository() *Repository {
	return &Repository{
		Fields: GlobFields,
	}
}

func (r *Repository) UpdateFieldImageID(ctx context.Context, imageID domain.ImageID, fieldID domain.FieldID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.Fields[fieldID]; !ok {
		return interr.NewNotFoundError("field does not exist")
	}

	r.Fields[fieldID] = &imageID

	return nil
}

func (r *Repository) ClearFieldImageID(ctx context.Context, fieldID domain.FieldID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.Fields[fieldID]; !ok {
		return interr.NewNotFoundError("field does not exist")
	}

	r.Fields[fieldID] = nil

	return nil
}
func (r *Repository) GetFieldImageID(ctx context.Context, fieldID domain.FieldID) (*domain.ImageID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.Fields[fieldID]; !ok {
		return nil, interr.NewNotFoundError("field does not exist")
	}

	return r.Fields[fieldID], nil
}
