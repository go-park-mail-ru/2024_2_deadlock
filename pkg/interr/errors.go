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
	Causes() interface{}
	Unwrap() error
}

type InternalError struct {
	ErrError  error
	ErrCauses interface{}
}

func (e InternalError) Error() string {
	return fmt.Sprintf("%s: causes: %v", e.ErrError, e.ErrCauses)
}

func (e InternalError) Causes() interface{} {
	return e.ErrCauses
}

func (e InternalError) Unwrap() error {
	return e.ErrError
}

func NewInternalError(err error, causes interface{}) InternalErr {
	return InternalError{
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) InternalErr {
	return InternalError{
		ErrError:  ErrNotFound,
		ErrCauses: causes,
	}
}

func NewAlreadyExistsError(causes interface{}) InternalErr {
	return InternalError{
		ErrError:  ErrAlreadyExists,
		ErrCauses: causes,
	}
}
