package http

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/utils"
)

type ResponseBody struct {
	Body interface{} `json:"body"`
}

func (s *Server) SendError(w http.ResponseWriter, restErr resterr.RestErr) {
	w.WriteHeader(restErr.Status())

	if err := utils.EncodeBody(w, restErr); err != nil {
		s.log.Errorw("could not encode error response", zap.Error(err))
	}
}

func (s *Server) ProcessBadRequestError(w http.ResponseWriter, err error) {
	restErr := resterr.NewBadRequestError(err)
	s.log.Errorw("could not decode user input", zap.Error(restErr))
	s.SendError(w, restErr)
}

func (s *Server) ProcessInternalServerError(w http.ResponseWriter, err error) {
	s.log.Errorw("internal server error", zap.Error(resterr.NewInternalServerError(err)))
	s.SendError(w, resterr.NewInternalServerError("internal server error"))
}

func (s *Server) SendBody(w http.ResponseWriter, v interface{}) {
	body := new(ResponseBody)
	body.Body = v

	err := utils.EncodeBody(w, body)
	if err != nil {
		s.ProcessInternalServerError(w, err)
	}
}
