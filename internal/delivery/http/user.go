package http

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/utils"
)

type UserUC interface {
	CurrentUser(ctx context.Context, userID domain.UserID) (domain.User, error)
}

func (s *Server) CurrentUser(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetCtxUserID(r.Context())
	if userID == 0 {
		utils.SendError(s.log, w, resterr.NewUnauthorizedError("unauthorized, please login"))
		return
	}

	user, err := s.uc.User.CurrentUser(r.Context(), userID)

	if errors.Is(err, resterr.ErrNotFound) {
		s.log.Errorw("could not get current user", zap.Error(err))
		utils.SendError(s.log, w, resterr.NewNotFoundError("user not found"))

		return
	}

	if err != nil {
		utils.ProcessInternalServerError(s.log, w, err)
		return
	}

	utils.SendBody(s.log, w, user)
}
