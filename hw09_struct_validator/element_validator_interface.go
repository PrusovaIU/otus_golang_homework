package hw09structvalidator

import (
	"reflect"
)

type ElementValidatorInterface interface {
	Validate(reflect.Value, reflect.Kind, string, string) error
}
