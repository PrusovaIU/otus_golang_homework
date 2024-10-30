package main

import (
	"errors"
	"fmt"
	"reflect"
)

func validateString(fieldValue reflect.Value, tag string) error {
	return nil
}

func validateInt(fieldValue reflect.Value, tag string) error {
	return nil
}

func validateSlice(fieldValue reflect.Value, fieldType reflect.StructField, tag string) ValidationErrors {
	var errs ValidationErrors = []ValidationError{}
	var validateFunc func(reflect.Value, string) error = nil
	switch fieldType.Type.Elem().Kind() {
	case reflect.Int:
		validateFunc = validateInt
	case reflect.String:
		validateFunc = validateString
	}
	if validateFunc != nil {
		for i := 0; i < fieldValue.Len(); i++ {
			elValue := fieldValue.Index(i)
			if err := validateFunc(elValue, tag); err != nil {
				validationErr := ValidationError{
					Field: fmt.Sprintf("%s[%d]", fieldType.Name, i),
					Err:   err,
				}
				errs = append(errs, validationErr)
			}
		}
	}
	return errs
}

func validateNotSlice(fieldValue reflect.Value, fieldType reflect.Kind, fieldName string, tag string) ValidationError {
	var err error = nil
	var validationErr = ValidationError{}
	switch fieldType {
	case reflect.String:
		err = validateString(fieldValue, tag)
	case reflect.Int:
		err = validateInt(fieldValue, tag)
	}
	if err != nil {
		validationErr = ValidationError{
			Field: fieldName,
			Err:   err,
		}
	}
	return validationErr
}

func validateField(fieldValue reflect.Value, fieldType reflect.StructField) ValidationErrors {
	isExported := fieldType.IsExported()
	var errs ValidationErrors = []ValidationError{}
	if isExported {
		tag := fieldType.Tag.Get("validate")
		if len(tag) > 0 {
			fieldKind := fieldType.Type.Kind()
			if fieldKind == reflect.Slice {
				errs = append(errs, validateSlice(fieldValue, fieldType, tag)...)
			} else {
				err := validateNotSlice(fieldValue, fieldKind, fieldType.Name, tag)
				errs = append(errs, err)
			}
		}
	}
	return errs
}

func Validate() error {
	var errs ValidationErrors = []ValidationError{}
	errs = append(errs, ValidationError{
		Field: "field_1",
		Err:   errors.New("test_error_1"),
	})
	errs = append(errs, ValidationError{
		Field: "field_2",
		Err:   errors.New("test_error_2"),
	})
	return errs
}

func main() {
	a := Validate()
	fmt.Println(a)
}
