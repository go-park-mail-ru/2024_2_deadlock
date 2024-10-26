package app

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http/common"
	v1 "github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http/v1"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
	imagerepo "github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/image"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/local/session"
	pgarticle "github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg/article"
	pguser "github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg/user"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/article"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/auth"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/avatar"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/fieldimage"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase/user"
)

type APIEntrypoint struct {
	Config *bootstrap.Config
	server *common.Server
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

	minioAdapter := adapters.NewMinioAdapter(e.Config)
	if err := minioAdapter.Init(); err != nil {
		logger.Errorw("init minio adapter error", zap.Error(err))
		return errors.Wrap(err, "init minio adapter")
	}

	userRepo := pguser.NewRepository(pgAdapter)
	sessionRepo := session.NewStorage()
	articleRepo := pgarticle.NewRepository(pgAdapter)
	avatarRepo := imagerepo.NewRepository(minioAdapter)
	fieldImageRepo := imagerepo.NewRepository(minioAdapter)

	if err := avatarRepo.Init(ctx, "avatar_bucket"); err != nil {
		logger.Errorw("init avatar repo error", zap.Error(err))
		return errors.Wrap(err, "init avatar repo")
	}

	if err := fieldImageRepo.Init(ctx, "field_image_bucket"); err != nil {
		logger.Errorw("init fieldImage repo error", zap.Error(err))
		return errors.Wrap(err, "init fieldImage repo")
	}

	authUC := auth.NewUsecase(auth.Repositories{
		Session: sessionRepo,
		User:    userRepo,
	})
	userUC := user.NewUsecase(user.Repositories{
		User:  userRepo,
		Image: avatarRepo,
	})
	articleUC := article.NewUsecase(article.Repositories{
		Article: articleRepo,
	})
	avatarUC := avatar.NewUsecase(avatar.Repositories{
		ImageRepo: avatarRepo,
		UserRepo:  userRepo,
	})
	fieldImageUC := fieldimage.NewUsecase(fieldimage.Repositories{
		ImageRepo: fieldImageRepo,
		// FieldRepo: fieldRepo,
	})

	ucs := v1.UseCases{
		User:       userUC,
		Auth:       authUC,
		Article:    articleUC,
		Avatar:     avatarUC,
		FieldImage: fieldImageUC,
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
