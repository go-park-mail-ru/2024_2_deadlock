package utils

import (
	"encoding/json"
	"net/http"
)

func EncodeBody(w http.ResponseWriter, v interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}

func DecodeBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
