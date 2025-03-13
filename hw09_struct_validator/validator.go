package hw09structvalidator

import validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"

func Validate(v interface{}) (validationErrs.ValidationErrors, error) {
	return NewValidator().Validate(v)
}
