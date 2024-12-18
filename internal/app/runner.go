package app

import (
	"context"
	"fmt"
	"io"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
)

type Entrypoint interface {
	io.Closer
	Init(ctx context.Context) error
	Run(ctx context.Context) error
}

func Run(ctx context.Context, e Entrypoint) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dg := depgraph.NewDepGraph()

	logger, err := dg.GetLogger()
	if err != nil {
		return fmt.Errorf("app.Run dg.GetLogger: %w", err)
	}

	if err := e.Init(ctx); err != nil {
		logger.Errorw("entrypoint init error", zap.Error(err))
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		logger.Info("starting app...")
		return e.Run(ctx)
	})

	// graceful shutdown
	eg.Go(func() error {
		<-ctx.Done()
		logger.Info("gracefully shutting down app...")

		return e.Close()
	})

	if err := eg.Wait(); err != nil {
		logger.Infof("app was shut down, reason: %s", err)
	}

	return nil
}
