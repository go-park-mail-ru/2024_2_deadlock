package v1

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

type AuthUC interface {
	Login(ctx context.Context, user *domain.UserInput) (domain.SessionID, error)
	Logout(ctx context.Context, sessionID domain.SessionID) error
	Register(ctx context.Context, user *domain.UserInput) (domain.SessionID, error)
	GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	input := new(domain.UserInput)

	err := utils.DecodeBody(r, input)
	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	sessionID, err := h.UC.Auth.Login(r.Context(), input)

	if errors.Is(err, interr.ErrNotFound) {
		h.log.Errorw("user not found", zap.Error(err))
		utils.SendError(h.log, w, resterr.NewNotFoundError("user not found"))

		return
	}

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	cfg := h.cfg.Server.Session.Cookie
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

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	cookie, err := r.Cookie("session_id")
	if err != nil {
		utils.SendError(h.log, w, resterr.NewUnauthorizedError("unauthorized, please login"))
		return
	}

	err = h.UC.Auth.Logout(r.Context(), domain.SessionID(cookie.Value))
	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	cfg := h.cfg.Server.Session.Cookie
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

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	input := new(domain.UserInput)

	err := utils.DecodeBody(r, input)
	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	sessionID, err := h.UC.Auth.Register(r.Context(), input)

	if errors.Is(err, interr.ErrAlreadyExists) {
		h.log.Errorw("user already exists", zap.Error(err))
		utils.SendError(h.log, w, resterr.NewConflictError("user already exists"))

		return
	}

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	cfg := h.cfg.Server.Session.Cookie
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
