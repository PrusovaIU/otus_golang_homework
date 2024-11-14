package main

import "reflect"

type SliceValidatorInterface interface {
	Validate(reflect.Value, reflect.StructField, string) ValidationErrors
}

type FieldFalidator struct {
	ElementValidator ElementValidatorInterface
	SliceValidator   SliceValidatorInterface
}

func NewFieldValidator() FieldFalidator {
	fv := FieldFalidator{}
	fv.ElementValidator = NewElementValidator()
	fv.SliceValidator = NewSliceValidator()
	return fv
}

func (fv FieldFalidator) Validate(fieldValue reflect.Value, fieldType reflect.StructField) ValidationErrors {
	isExported := fieldType.IsExported()
	var errs ValidationErrors = []ValidationError{}
	if isExported {
		tag := fieldType.Tag.Get("validate")
		if len(tag) > 0 {
			fieldKind := fieldType.Type.Kind()
			if fieldKind == reflect.Slice {
				errs = append(errs, fv.SliceValidator.Validate(fieldValue, fieldType, tag)...)
			} else {
				err := fv.ElementValidator.Validate(fieldValue, fieldKind, fieldType.Name, tag)
				errs = append(errs, err)
			}
		}
	}
	return errs
}
