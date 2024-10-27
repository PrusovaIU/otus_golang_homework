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

func validateSlice(fieldValue reflect.Value, fieldType reflect.StructField, tag string) []error {
	var errs []error = []error{}
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
				errs = append(errs, err)
			}
		}
	}

}

func validateField(fieldValue reflect.Value, fieldType reflect.StructField) error {
	isExported := fieldType.IsExported()
	var err error = nil
	if isExported {
		tag := fieldType.Tag.Get("validate")
		if len(tag) > 0 {
			switch fieldType.Type.Kind() {
			case reflect.String:
				err = validateString(fieldValue, tag)
			case reflect.Int:
				err = validateInt(fieldValue, tag)
			case reflect.Slice:
				err = validateSlice(fieldValue, fieldType, tag)
			}
		}
	}
	return err
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
