package resterr

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest     = errors.New("bad request")
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
	ErrForbidden      = errors.New("forbidden")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInternalServer = errors.New("internal server error")
)

type RestErr interface {
	Status() int
	Error() string
	Unwrap() error
}

type RestError struct {
	StatusCode int         `json:"-"`
	Err        error       `json:"-"`
	Code       string      `json:"code,omitempty"`
	Causes     interface{} `json:"message,omitempty"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("%s: %s", e.Causes, e.Err)
}

func (e RestError) Unwrap() error {
	return e.Err
}

func (e RestError) Status() int {
	return e.StatusCode
}

func NewRestError(status int, code string, err error, causes interface{}) RestErr {
	return RestError{
		StatusCode: status,
		Err:        err,
		Code:       code,
		Causes:     causes,
	}
}

func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		StatusCode: http.StatusBadRequest,
		Err:        ErrBadRequest,
		Code:       "bad_request",
		Causes:     causes,
	}
}

func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		StatusCode: http.StatusNotFound,
		Err:        ErrNotFound,
		Code:       "not_found",
		Causes:     causes,
	}
}

func NewConflictError(causes interface{}) RestErr {
	return RestError{
		StatusCode: http.StatusConflict,
		Err:        ErrConflict,
		Code:       "conflict",
		Causes:     causes,
	}
}

func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		StatusCode: http.StatusForbidden,
		Err:        ErrForbidden,
		Code:       "forbidden",
		Causes:     causes,
	}
}

func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		StatusCode: http.StatusUnauthorized,
		Err:        ErrUnauthorized,
		Code:       "unauthorized",
		Causes:     causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	return RestError{
		StatusCode: http.StatusInternalServerError,
		Err:        ErrInternalServer,
		Code:       "internal_server_error",
		Causes:     causes,
	}
}
