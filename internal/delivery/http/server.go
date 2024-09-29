package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
)

type UseCases struct {
	// TODO: add use cases
}

type Server struct {
	cfg  *bootstrap.Config
	log  *zap.SugaredLogger
	http *http.Server
	mux  *mux.Router
	uc   UseCases
}

func NewServer(cfg *bootstrap.Config, log *zap.SugaredLogger, uc UseCases) *Server {
	s := &Server{
		cfg: cfg,
		log: log,
		uc:  uc,
	}
	s.makeRoutes()

	return s
}

func (s *Server) Run() error {
	s.http = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", s.cfg.Server.Port),
		Handler: s.mux,
	}

	if err := s.http.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		return errors.Wrap(err, "http server listen and serve")
	}

	return nil
}

func (s *Server) Close() error {
	if s.http != nil {
		err := s.http.Close()
		if err != nil {
			s.log.Errorf("could not stop http server: %v", err)
			return errors.Wrap(err, "http server close")
		}

		s.log.Info("http server stopped")
	}

	return nil
}
