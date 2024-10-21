package middleware

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	v1 "github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http/v1"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

func AuthMW(log *zap.SugaredLogger, cfg *bootstrap.Config, auth v1.AuthUC) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie(cfg.Server.Session.Cookie.Name)

			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}

			if err != nil {
				log.Errorw("auth mw r.Cookie get", zap.Error(err))
				utils.ProcessBadRequestError(log, w, err)

				return
			}

			id, err := auth.GetUserID(r.Context(), utils.GetCookieSessionID(cfg, r))
			if errors.Is(err, resterr.ErrNotFound) {
				return
			}

			if err != nil {
				log.Errorw("auth mw auth.GetUserID", zap.Error(err))
				utils.ProcessBadRequestError(log, w, err)

				return
			}

			ctx := context.WithValue(r.Context(), utils.CtxKeyUserID{}, id)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
