package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
)

type AuthRepository interface {
	GetUser(ctx context.Context, user *domain.UserInput) (*domain.User, error)
	GetUserByID(ctx context.Context, user domain.UserID) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.UserInput) (*domain.User, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, userID domain.UserID) (domain.SessionID, error)
	DeleteSession(ctx context.Context, sessionID domain.SessionID) error
	GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error)
}

type AuthRepositories struct {
	Auth    AuthRepository
	Session SessionRepository
}

type AuthUsecase struct {
	repo AuthRepositories
}

func NewAuthUsecase(repo AuthRepositories) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (uc *AuthUsecase) Login(ctx context.Context, user *domain.UserInput) (domain.SessionID, error) {
	u, err := uc.repo.Auth.GetUser(ctx, user)
	if err != nil {
		return "", err
	}

	return uc.repo.Session.CreateSession(ctx, u.ID)
}

func (uc *AuthUsecase) Logout(ctx context.Context, sessionID domain.SessionID) error {
	return uc.repo.Session.DeleteSession(ctx, sessionID)
}

func (uc *AuthUsecase) Register(ctx context.Context, user *domain.UserInput) (domain.SessionID, error) {
	u, err := uc.repo.Auth.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return uc.repo.Session.CreateSession(ctx, u.ID)
}

func (uc *AuthUsecase) GetUserID(ctx context.Context, sessionID domain.SessionID) (domain.UserID, error) {
	return uc.repo.Session.GetUserID(ctx, sessionID)
}

func (uc *AuthUsecase) CurrentUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	return uc.repo.Auth.GetUserByID(ctx, userID)
}
