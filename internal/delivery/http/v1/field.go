package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

type FieldUC interface {
	GetFieldInfo(ctx context.Context, fieldID domain.FieldID) (*domain.FieldInfo, error)
}

func (h *Handler) GetFieldInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fieldIDStr := vars["fieldID"]
	fieldIDInt, err := strconv.Atoi(fieldIDStr)

	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	fieldID := domain.FieldID(fieldIDInt)

	info, err := h.UC.Field.GetFieldInfo(r.Context(), fieldID)

	if errors.Is(err, interr.ErrNotFound) {
		if errors.Is(err, interr.ErrNotFound) {
			h.log.Errorw("field not found", zap.Error(err))
			utils.SendError(h.log, w, resterr.NewNotFoundError("field not found"))

			return
		}
	}

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, info)
}
