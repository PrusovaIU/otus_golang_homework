package hw09structvalidator

import (
	"reflect"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

type SliceValidatorInterface interface {
	Validate(reflect.Value, reflect.StructField, string) (validationErrs.ValidationErrors, error)
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

func (fv FieldValidator) checkType(
	fieldValue reflect.Value, fieldType reflect.StructField, tag string,
) (validationErrs.ValidationErrors, error) {
	var errs validationErrs.ValidationErrors = []validationErrs.ValidationError{}
	fieldKind := fieldType.Type.Kind()
	if fieldKind == reflect.Slice {
		validationErrs, err := fv.SliceValidator.Validate(fieldValue, fieldType, tag)
		if err == nil {
			errs = append(errs, validationErrs...)
		} else {
			return nil, err
		}
	} else {
		err := fv.ElementValidator.Validate(fieldValue, fieldKind, fieldType.Name, tag)
		if elValidatorErr, ok := err.(validationErrs.ValidationError); ok {
			errs = append(errs, elValidatorErr)
		} else if err != nil {
			return nil, err
		}
	}
	return errs, nil
}

// Validate проверяет значение поля структуры на соответствие заданному условию.
// Входные параметры:
// fieldValue - значение поля
// fieldType - тип поля
// Возвращаемое значение:
// errors.ValidationErrors - список ошибок валидации.
func (fv FieldValidator) Validate(fieldValue reflect.Value, fieldType reflect.StructField) (validationErrs.ValidationErrors, error) {
	var errs validationErrs.ValidationErrors = []validationErrs.ValidationError{}
	if isExported := fieldType.IsExported(); isExported {
		tag := fieldType.Tag.Get("validate")
		if len(tag) > 0 {
			return fv.checkType(fieldValue, fieldType, tag)
		}
	}
	return errs, nil
}
