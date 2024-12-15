package hw09structvalidator

import (
	"reflect"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

type SliceValidatorInterface interface {
	Validate(reflect.Value, reflect.StructField, string) errors.ValidationErrors
}

type FieldValidator struct {
	ElementValidator ElementValidatorInterface
	SliceValidator   SliceValidatorInterface
}

func NewFieldValidator() FieldValidator {
	fv := FieldValidator{}
	fv.ElementValidator = NewElementValidator()
	fv.SliceValidator = NewSliceValidator()
	return fv
}

// Validate проверяет значение поля структуры на соответствие заданному условию.
// Входные параметры:
// fieldValue - значение поля
// fieldType - тип поля
// Возвращаемое значение:
// errors.ValidationErrors - список ошибок валидации
func (fv FieldValidator) Validate(fieldValue reflect.Value, fieldType reflect.StructField) errors.ValidationErrors {
	var errs errors.ValidationErrors = []errors.ValidationError{}
	tag := fieldType.Tag.Get("validate")
	if len(tag) > 0 {
		fieldKind := fieldType.Type.Kind()
		if fieldKind == reflect.Slice {
			errs = append(errs, fv.SliceValidator.Validate(fieldValue, fieldType, tag)...)
		} else {
			err := fv.ElementValidator.Validate(fieldValue, fieldKind, fieldType.Name, tag)
			if err.IsErr() {
				errs = append(errs, err)
			}
		}
	}
	return errs
}
