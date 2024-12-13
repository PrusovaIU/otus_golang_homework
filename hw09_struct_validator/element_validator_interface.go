package hw09structvalidator

import (
	"reflect"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

type ElementValidatorInterface interface {
	Validate(reflect.Value, reflect.Kind, string, string) errors.ValidationError
}
