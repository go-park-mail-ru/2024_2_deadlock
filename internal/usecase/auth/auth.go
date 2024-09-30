package auth

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type Usecase struct {
	usecase.CommonUC
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
