package utils

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
)

func GetCookieSessionID(cfg *bootstrap.Config, r *http.Request) domain.SessionID {
	cookie, err := r.Cookie(cfg.Server.Session.Cookie.Name)
	if err != nil {
		return ""
	}

	return domain.SessionID(cookie.Value)
}
