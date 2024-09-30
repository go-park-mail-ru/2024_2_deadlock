package validator

import (
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
)

func ValidateStruct(s interface{}) error {
	dg := depgraph.NewDepGraph()

	validate, err := dg.GetValidator()
	if err != nil {
		return err
	}

	return validate.Struct(s)
}
