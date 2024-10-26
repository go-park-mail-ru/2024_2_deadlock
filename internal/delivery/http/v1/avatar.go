package v1

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

type AvatarUC interface {
	SetAvatarImage(ctx context.Context, data *domain.ImageData, userID domain.UserID) (domain.ImageURL, error)
	DeleteAvatarImage(ctx context.Context, userID domain.UserID) error
}

func (h *Handler) SetAvatarImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	userID := utils.GetCtxUserID(r.Context())
	if userID == 0 {
		utils.SendError(h.log, w, resterr.NewUnauthorizedError("unauthorized, please login"))
		return
	}

	imageData := new(domain.ImageData)

	err := utils.DecodeBody(r, imageData)
	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	imageURL, err := h.UC.Avatar.SetAvatarImage(r.Context(), imageData, userID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, imageURL)
}

func (h *Handler) DeleteAvatarImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	userID := utils.GetCtxUserID(r.Context())
	if userID == 0 {
		utils.SendError(h.log, w, resterr.NewUnauthorizedError("unauthorized, please login"))
		return
	}

	err := h.UC.Avatar.DeleteAvatarImage(r.Context(), userID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, struct{}{})
}
