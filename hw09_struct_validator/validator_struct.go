package main

import (
	"errors"
	"reflect"
)

type FieldValidatorInterface interface {
	Validate(reflect.Value, reflect.StructField) ValidationErrors
}

type Validator struct {
	FieldValidator FieldValidatorInterface
}

func NewValidator() Validator {
	v := Validator{}
	v.FieldValidator = NewFieldValidator()
	return v
}

func (v Validator) Validate(value interface{}) ValidationErrors {
	errs := []ValidationError{}

	vValue := reflect.ValueOf(v)
	vType := reflect.TypeOf(v)

	if vValue.Kind() != reflect.Struct {
		errs = append(errs, ValidationError{
			Field: "Root",
			Err:   errors.New("expected struct"),
		})
		return errs
	}

	for i := 0; i < vValue.NumField(); i++ {
		fieldValue := vValue.Field(i)
		fieldType := vType.Field(i)

		fieldErrs := v.FieldValidator.Validate(fieldValue, fieldType)
		errs = append(errs, fieldErrs...)
	}

	return errs
}
