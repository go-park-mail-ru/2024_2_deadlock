package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/errors"
)

type ResponseBody struct {
	Body interface{} `json:"body"`
}

func SendError(w http.ResponseWriter, err errors.RestErr) {
	w.WriteHeader(err.Status())
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(err)
}

func SendBody(w http.ResponseWriter, v interface{}) error {
	encoder := json.NewEncoder(w)

	body := new(ResponseBody)
	body.Body = v

	if err := encoder.Encode(body); err != nil {
		SendError(w, errors.NewInternalServerError(err))
		return err
	}

	return nil
}

func DecodeBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
