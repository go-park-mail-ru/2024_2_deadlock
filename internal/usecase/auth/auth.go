package auth

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, userID domain.UserID) (domain.SessionID, error)
	DeleteSession(ctx context.Context, sessionID domain.SessionID) error
	GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, userID domain.UserID) (domain.User, error)
	Create(ctx context.Context, user domain.UserInput) (domain.UserID, error)
}

type Repositories struct {
	Session SessionRepository
	User    UserRepository
}

type Usecase struct {
	usecase.CommonUC
	repo Repositories
}

func (uc *Usecase) Login(ctx context.Context, user domain.UserInput) (domain.SessionID, error) {
	panic("implement me")
}

func (uc *Usecase) Logout(ctx context.Context, sessionID domain.SessionID) error {
	panic("implement me")
}

func (uc *Usecase) Register(ctx context.Context, user domain.UserInput) (domain.SessionID, error) {
	panic("implement me")
}
