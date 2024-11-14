package main

import "reflect"

type ElementValidatorInterface interface {
	Validate(reflect.Value, reflect.Kind, string, string) ValidationError
}
