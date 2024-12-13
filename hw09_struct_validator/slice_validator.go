package hw09structvalidator

import (
	"fmt"
	"reflect"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

type SliceValidator struct {
	ElementValidator ElementValidatorInterface
}

func NewSliceValidator() SliceValidator {
	sv := SliceValidator{}
	sv.ElementValidator = NewElementValidator()
	return sv
}

func (sv SliceValidator) Validate(fieldValue reflect.Value, fieldType reflect.StructField, tag string) errors.ValidationErrors {
	var errs errors.ValidationErrors = []errors.ValidationError{}
	for i := 0; i < fieldValue.Len(); i++ {
		elValue := fieldValue.Index(i)
		err := sv.ElementValidator.Validate(elValue, fieldType.Type.Elem().Kind(), fmt.Sprintf("%s[%d]", fieldType.Name, i), tag)
		errs = append(errs, err)
	}
	return errs
}
