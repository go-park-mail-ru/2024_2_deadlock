package auth

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type SessionRepository interface {
	Create(ctx context.Context, userID domain.UserID) (domain.SessionID, error)
	Delete(ctx context.Context, sessionID domain.SessionID) error
	GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error)
}

type UserRepository interface {
	Get(ctx context.Context, user *domain.UserInput) (*domain.User, error)
	Create(ctx context.Context, user *domain.UserInput) (*domain.User, error)
}

type Repositories struct {
	Session SessionRepository
	User    UserRepository
}

type Usecase struct {
	usecase.CommonUC
	repo Repositories
}

func (uc *Usecase) Login(ctx context.Context, user *domain.UserInput) (domain.SessionID, error) {
	u, err := uc.repo.User.Get(ctx, user)
	if err != nil {
		return "", err
	}

	return uc.repo.Session.Create(ctx, u.ID)
}

func (uc *Usecase) Logout(ctx context.Context, sessionID domain.SessionID) error {
	return uc.repo.Session.Delete(ctx, sessionID)
}

func (uc *Usecase) Register(ctx context.Context, user *domain.UserInput) (domain.SessionID, error) {
	u, err := uc.repo.User.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return uc.repo.Session.Create(ctx, u.ID)
}
