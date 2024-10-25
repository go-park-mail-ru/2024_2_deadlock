package v1

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
)

type ImageUC interface {
	SetFieldPhoto(ctx context.Context, data *domain.ImageData) (domain.ImageURL, error)
	GetFieldPhoto(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error)
	UpdateFieldPhoto(ctx context.Context, data *domain.ImageData, imageID domain.ImageID) (domain.ImageURL, error)
	DeleteFieldPhoto(ctx context.Context, imageID domain.ImageID) error
}

// func (h *Handler) SetFieldImage(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	image, err := h.UC.Image.SetFieldPhoto(r.Context())
// }
