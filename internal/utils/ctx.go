package utils

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
)

type CtxKeyUserID struct{}

func GetCtxUserID(ctx context.Context) domain.UserID {
	userID, ok := ctx.Value(CtxKeyUserID{}).(domain.UserID)
	if !ok {
		return 0
	}

	return userID
}
