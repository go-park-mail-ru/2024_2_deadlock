package app

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/local/session"
	pguser "github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg/user"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/auth"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/user"
)

type APIEntrypoint struct {
	Config *bootstrap.Config
	server *http.Server
}

func (e *APIEntrypoint) Init(ctx context.Context) error {
	dg := depgraph.NewDepGraph()

	logger, err := dg.GetLogger()
	if err != nil {
		return errors.Wrap(err, "get logger")
	}

	pgAdapter := adapters.NewAdapterPG(e.Config)
	if err := pgAdapter.Init(ctx); err != nil {
		logger.Errorw("init pg adapter error", zap.Error(err))
		return errors.Wrap(err, "init pg adapter")
	}

	userRepo := pguser.NewRepository(pgAdapter)
	sessionRepo := session.NewStorage()

	authUC := auth.NewUsecase(auth.Repositories{
		Session: sessionRepo,
		User:    userRepo,
	})
	userUC := user.NewUsecase(user.Repositories{
		User: userRepo,
	})

	ucs := http.UseCases{
		User: userUC,
		Auth: authUC,
	}

	e.server = http.NewServer(e.Config, logger, ucs)

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
