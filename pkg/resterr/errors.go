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
	Causes() interface{}
	Unwrap() error
}

type RestError struct {
	ErrStatus int         `json:"-"`
	ErrError  error       `json:"-"`
	ErrCode   string      `json:"code,omitempty"`
	ErrCauses interface{} `json:"message,omitempty"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - error: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func (e RestError) Unwrap() error {
	return e.ErrError
}

func NewRestError(status int, code string, err error, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCode:   code,
		ErrCauses: causes,
	}
}

func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  ErrBadRequest,
		ErrCode:   "bad_request",
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  ErrNotFound,
		ErrCode:   "not_found",
		ErrCauses: causes,
	}
}

func NewConflictError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusConflict,
		ErrError:  ErrConflict,
		ErrCode:   "conflict",
		ErrCauses: causes,
	}
}

func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  ErrForbidden,
		ErrCode:   "forbidden",
		ErrCauses: causes,
	}
}

func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  ErrUnauthorized,
		ErrCode:   "unauthorized",
		ErrCauses: causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  ErrInternalServer,
		ErrCode:   "internal_server_error",
		ErrCauses: causes,
	}
}
