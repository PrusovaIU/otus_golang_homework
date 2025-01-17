package hw09structvalidator

import (
	"errors"
	"reflect"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

type FieldValidatorInterface interface {
	Validate(reflect.Value, reflect.StructField) (validationErrs.ValidationErrors, error)
}

type Validator struct {
	FieldValidator FieldValidatorInterface
}

func NewValidator() Validator {
	v := Validator{}
	v.FieldValidator = NewFieldValidator()
	return v
}

func (v Validator) Validate(value interface{}) (validationErrs.ValidationErrors, error) {
	errs := []validationErrs.ValidationError{}

	vValue := reflect.ValueOf(value)
	vType := reflect.TypeOf(value)

	if vValue.Kind() != reflect.Struct {
		errs = append(errs, validationErrs.ValidationError{
			Field: "Root",
			Err:   errors.New("expected struct"),
		})
		return errs, nil
	}

	for i := 0; i < vValue.NumField(); i++ {
		fieldValue := vValue.Field(i)
		fieldType := vType.Field(i)

		fieldErrs, err := v.FieldValidator.Validate(fieldValue, fieldType)
		if err != nil {
			return nil, err
		}
		errs = append(errs, fieldErrs...)
	}

	return errs, nil
}
