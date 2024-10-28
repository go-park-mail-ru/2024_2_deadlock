package app

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http/common"
	v1 "github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http/v1"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/local"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type APIEntrypoint struct {
	Config *bootstrap.Config
	server *common.Server
}

func (e *APIEntrypoint) Init(ctx context.Context) error {
	dg := depgraph.NewDepGraph()

	logger, err := dg.GetLogger()
	if err != nil {
		return fmt.Errorf("APIEntrypoint.Init dg.GetLogger: %w", err)
	}

	pgAdapter := adapters.NewAdapterPG(e.Config)
	if err := pgAdapter.Init(ctx); err != nil {
		logger.Errorw("init pg adapter error", zap.Error(err))
		return fmt.Errorf("APIEntrypoint.Init pgAdapter.Init: %w", err)
	}

	authRepo := pg.NewAuthRepository(pgAdapter)
	sessionRepo := local.NewSessionRepository()
	articleRepo := pg.NewArticleRepository(pgAdapter)

	authUC := usecase.NewAuthUsecase(usecase.AuthRepositories{
		Auth:    authRepo,
		Session: sessionRepo,
	})
	articleUC := usecase.NewArticleUsecase(usecase.ArticleRepositories{
		Article: articleRepo,
	})

	ucs := v1.UseCases{
		Auth:    authUC,
		Article: articleUC,
	}

	handlerV1 := v1.NewHandler(e.Config, logger, ucs)
	handlers := common.Handlers{V1: handlerV1}

	e.server = common.NewServer(e.Config, logger, handlers)

	return nil
}

func (e *APIEntrypoint) Run(ctx context.Context) error {
	return e.server.Run()
}

func (e *APIEntrypoint) Close() error {
	if e.server != nil {
		return e.server.Close()
	}

	return nil
}
