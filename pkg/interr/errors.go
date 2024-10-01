package interr

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type InternalErr interface {
	Error() string
	Unwrap() error
}

type InternalError struct {
	error  error
	causes interface{}
}

func (e InternalError) Error() string {
	return fmt.Sprintf("%s: %s", e.causes, e.error)
}

func (e InternalError) Unwrap() error {
	return e.error
}

func NewInternalError(err error, causes interface{}) InternalErr {
	return InternalError{
		error:  err,
		causes: causes,
	}
}

func NewNotFoundError(causes interface{}) InternalErr {
	return InternalError{
		error:  ErrNotFound,
		causes: causes,
	}
}

func NewAlreadyExistsError(causes interface{}) InternalErr {
	return InternalError{
		error:  ErrAlreadyExists,
		causes: causes,
	}
}
