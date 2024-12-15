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

// Validate проверяет каждый элементы слайса на соответствие заданным условиям.
// Входные параметры:
// fieldValue - значение поля структуры
// fieldType - тип поля структуры
// tag - тег валидации
// Возвращаемое значение:
// errs - список ошибок валидацииS
func (sv SliceValidator) Validate(fieldValue reflect.Value, fieldType reflect.StructField, tag string) errors.ValidationErrors {
	var errs errors.ValidationErrors = []errors.ValidationError{}
	for i := 0; i < fieldValue.Len(); i++ {
		elValue := fieldValue.Index(i)
		err := sv.ElementValidator.Validate(elValue, fieldType.Type.Elem().Kind(), fmt.Sprintf("%s[%d]", fieldType.Name, i), tag)
		if err.IsErr() {
			errs = append(errs, err)
		}
	}
	return errs
}
