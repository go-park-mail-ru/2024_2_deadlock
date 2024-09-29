package app

import (
	"context"

	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
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

	e.server = http.NewServer(e.Config, logger, http.UseCases{})

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
