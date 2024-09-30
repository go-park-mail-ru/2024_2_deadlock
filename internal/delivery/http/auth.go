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

type AuthUC interface {
	Login(ctx context.Context, user domain.UserInput) (domain.SessionID, error)
	Logout(ctx context.Context, sessionID domain.SessionID) error
	Register(ctx context.Context, user domain.UserInput) (domain.SessionID, error)
	CurrentUser(ctx context.Context, sessionID domain.SessionID) (domain.User, error)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	input := new(domain.UserInput)

	err := utils.DecodeBody(r, input)
	if err != nil {
		s.ProcessBadRequestError(w, err)
		return
	}

	sessionID, err := s.uc.Auth.Login(r.Context(), *input)

	if errors.Is(err, resterr.ErrNotFound) {
		s.log.Errorw("could not login", zap.Error(err))
		s.SendError(w, resterr.NewNotFoundError("user not found"))

		return
	}

	if err != nil {
		s.ProcessInternalServerError(w, err)
		return
	}

	cfg := s.cfg.Server.Session.Cookie
	cookie := &http.Cookie{
		Name:     cfg.Name,
		Value:    string(sessionID),
		Path:     cfg.Path,
		MaxAge:   int(cfg.MaxAge.Seconds()),
		HttpOnly: cfg.HttpOnly,
		Secure:   cfg.Secure,
	}

	http.SetCookie(w, cookie)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	cookie, err := r.Cookie("session_id")
	if err != nil {
		s.log.Errorw("could not get session id from cookies", zap.Error(err))
		s.SendError(w, resterr.NewUnauthorizedError(err))

		return
	}

	err = s.uc.Auth.Logout(r.Context(), domain.SessionID(cookie.Value))
	if err != nil {
		s.ProcessInternalServerError(w, err)
		return
	}

	cfg := s.cfg.Server.Session.Cookie
	cookie = &http.Cookie{
		Name:     cfg.Name,
		Value:    "",
		Path:     cfg.Path,
		MaxAge:   0,
		HttpOnly: cfg.HttpOnly,
		Secure:   cfg.Secure,
	}

	http.SetCookie(w, cookie)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	input := new(domain.UserInput)

	err := utils.DecodeBody(r, input)
	if err != nil {
		s.ProcessBadRequestError(w, err)
		return
	}

	sessionID, err := s.uc.Auth.Register(r.Context(), *input)

	if errors.Is(err, resterr.ErrConflict) {
		s.log.Errorw("could not register", zap.Error(err))
		s.SendError(w, resterr.NewConflictError("user already exists"))

		return
	}

	if err != nil {
		s.ProcessInternalServerError(w, err)
		return
	}

	cfg := s.cfg.Server.Session.Cookie
	cookie := &http.Cookie{
		Name:     cfg.Name,
		Value:    string(sessionID),
		Path:     cfg.Path,
		MaxAge:   int(cfg.MaxAge.Seconds()),
		HttpOnly: cfg.HttpOnly,
		Secure:   cfg.Secure,
	}

	http.SetCookie(w, cookie)
}

func (s *Server) CurrentUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sessionID, err := r.Cookie(s.cfg.Server.Session.Cookie.Name)
	if err != nil {
		s.log.Errorw("could not get session id from cookies", zap.Error(err))
		s.SendError(w, resterr.NewUnauthorizedError(err))

		return
	}

	user, err := s.uc.Auth.CurrentUser(r.Context(), domain.SessionID(sessionID.Value))

	if errors.Is(err, resterr.ErrNotFound) {
		s.log.Errorw("could not get current user", zap.Error(err))
		s.SendError(w, resterr.NewNotFoundError("user not found"))

		return
	}

	if err != nil {
		s.ProcessInternalServerError(w, err)
		return
	}

	s.SendBody(w, user)
}
