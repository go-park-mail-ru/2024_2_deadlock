package utils

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

type ResponseBody struct {
	Body interface{} `json:"body"`
}

func SendError(log *zap.SugaredLogger, w http.ResponseWriter, restErr resterr.RestErr) {
	w.WriteHeader(restErr.Status())

	if err := EncodeBody(w, restErr); err != nil {
		log.Errorw("could not encode error response", zap.Error(err))
	}
}

func ProcessBadRequestError(log *zap.SugaredLogger, w http.ResponseWriter, err error) {
	restErr := resterr.NewBadRequestError(err)
	log.Errorw("could not decode user input", zap.Error(restErr))
	SendError(log, w, restErr)
}

func ProcessInternalServerError(log *zap.SugaredLogger, w http.ResponseWriter, err error) {
	log.Errorw("internal server error", zap.Error(resterr.NewInternalServerError(err)))
	SendError(log, w, resterr.NewInternalServerError("internal server error"))
}

func SendBody(log *zap.SugaredLogger, w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	body := new(ResponseBody)
	body.Body = v

	err := EncodeBody(w, body)
	if err != nil {
		ProcessInternalServerError(log, w, err)
	}
}
