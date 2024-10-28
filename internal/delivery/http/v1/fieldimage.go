package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
)

type FieldImageUC interface {
	SetFieldImage(ctx context.Context, data *domain.ImageData, fieldID domain.FieldID) (domain.ImageURL, error)
	DeleteFieldImage(ctx context.Context, fieldID domain.FieldID) error
}

func (h *Handler) SetFieldImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	fieldIDStr := vars["fieldID"]
	fieldIDInt, err := strconv.Atoi(fieldIDStr)

	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	fieldID := domain.FieldID(fieldIDInt)

	imageData := new(domain.ImageData)

	file, fileHeader, multipartErr := utils.DecodeImage(r, h.cfg)
	if multipartErr != nil {
		h.log.Errorw("problems when image decoded from multipart", zap.Error(multipartErr))
		utils.SendError(h.log, w, multipartErr)

		return
	}

	imageData.Image = file
	imageData.Header = fileHeader

	imageURL, err := h.UC.FieldImage.SetFieldImage(r.Context(), imageData, fieldID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)

		return
	}

	utils.SendBody(h.log, w, imageURL)
}

func (h *Handler) DeleteFieldImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	fieldIDStr := vars["fieldID"]
	fieldIDInt, err := strconv.Atoi(fieldIDStr)

	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	fieldID := domain.FieldID(fieldIDInt)

	err = h.UC.FieldImage.DeleteFieldImage(r.Context(), fieldID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, struct{}{})
}
