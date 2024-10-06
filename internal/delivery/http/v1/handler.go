package v1

import (
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
)

type UseCases struct {
	Auth    AuthUC
	User    UserUC
	Article ArticleUC
}

type Handler struct {
	cfg *bootstrap.Config
	log *zap.SugaredLogger
	UC  UseCases
}

func NewHandler(cfg *bootstrap.Config, log *zap.SugaredLogger, uc UseCases) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
		UC:  uc,
	}
}
