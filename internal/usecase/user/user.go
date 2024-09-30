package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type Usecase struct {
	usecase.CommonUC
}

func (uc *Usecase) CurrentUser(ctx context.Context, userID domain.UserID) (domain.User, error) {
	panic("implement me")
}
