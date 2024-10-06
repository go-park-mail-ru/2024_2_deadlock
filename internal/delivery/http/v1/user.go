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

type UserUC interface {
	CurrentUser(ctx context.Context, userID domain.UserID) (*domain.User, error)
	GetUserInfo(ctx context.Context, userID domain.UserID) (*domain.UserInfo, error)
	UpdateUserInfo(ctx context.Context, userInfo *domain.UserUpdate, userID domain.UserID) error
}

func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetCtxUserID(r.Context())
	if userID == 0 {
		utils.SendError(h.log, w, resterr.NewUnauthorizedError("unauthorized, please login"))
		return
	}

	user, err := h.UC.User.CurrentUser(r.Context(), userID)

	if errors.Is(err, interr.ErrNotFound) {
		h.log.Errorw("current user not found", zap.Error(err))
		utils.SendError(h.log, w, resterr.NewNotFoundError("user not found"))

		return
	}

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, user)
}

func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	userIDInt, err := strconv.Atoi(userIDStr)

	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	userID := domain.UserID(userIDInt)

	info, err := h.UC.User.GetUserInfo(r.Context(), userID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, info)
}

func (h *Handler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	userID := utils.GetCtxUserID(r.Context())
	if userID == 0 {
		utils.SendError(h.log, w, resterr.NewUnauthorizedError("unauthorized, please login"))
		return
	}

	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	userIDInt, err := strconv.Atoi(userIDStr)

	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	userIDInput := domain.UserID(userIDInt)
	if userID != userIDInput {
		utils.SendError(h.log, w, resterr.NewForbiddenError("no rights to change this user"))
	}

	updateData := new(domain.UserUpdate)

	err = utils.DecodeBody(r, updateData)
	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	err = h.UC.User.UpdateUserInfo(r.Context(), updateData, userID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, struct{}{})
}
