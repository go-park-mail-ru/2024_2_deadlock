package utils

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
)

func GetCookieSessionID(r *http.Request, cfg *bootstrap.Config) domain.SessionID {
	cookie, err := r.Cookie(cfg.Server.Session.Cookie.Name)
	if err != nil {
		return ""
	}

	return domain.SessionID(cookie.Value)
}

func SetCookieSession(w http.ResponseWriter, cfg *bootstrap.Config, sessionID string) {
	cookieCfg := cfg.Server.Session.Cookie
	cookie := &http.Cookie{
		Name:     cookieCfg.Name,
		Value:    sessionID,
		Path:     cookieCfg.Path,
		MaxAge:   int(cookieCfg.MaxAge.Seconds()),
		HttpOnly: cookieCfg.HttpOnly,
		Secure:   cookieCfg.Secure,
	}

	http.SetCookie(w, cookie)
}

func DeleteCookieSession(w http.ResponseWriter, cfg *bootstrap.Config) {
	cookieCfg := cfg.Server.Session.Cookie
	cookie := &http.Cookie{
		Name:     cookieCfg.Name,
		Value:    "",
		Path:     cookieCfg.Path,
		MaxAge:   0,
		HttpOnly: cookieCfg.HttpOnly,
		Secure:   cookieCfg.Secure,
	}

	http.SetCookie(w, cookie)
}
