package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func validateString(fieldValue reflect.Value, cond_name string, cond_valueL string) error {
	return nil
}

func validateInt(fieldValue reflect.Value, cond_name string, cond_valueL string) error {
	return nil
}

func validateSlice(fieldValue reflect.Value, fieldType reflect.StructField, tag string) ValidationErrors {
	var errs ValidationErrors = []ValidationError{}
	for i := 0; i < fieldValue.Len(); i++ {
		elValue := fieldValue.Index(i)
		err := validateNotSlice(elValue, fieldType.Type.Elem().Kind(), fmt.Sprintf("%s[%d]", fieldType.Name, i), tag)
		errs = append(errs, err)
	}
	return errs
}

func parseTag(tag string) (string, string, error) {
	split := strings.Split(tag, ":")
	if len(split) != 2 {
		return "", "", errors.New("wrong tag format")
	}
	condition := strings.Trim(split[0], " ")
	value := strings.Trim(split[1], " ")
	return condition, value, nil
}

func validateNotSlice(fieldValue reflect.Value, fieldType reflect.Kind, fieldName string, tag string) ValidationError {
	var err error = nil
	var validationErr = ValidationError{}
	condition, condition_value, err := parseTag(tag)
	if err == nil {
		switch fieldType {
		case reflect.String:
			err = validateString(fieldValue, condition, condition_value)
		case reflect.Int:
			err = validateInt(fieldValue, condition, condition_value)
		}
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

func Validate(v interface{}) ValidationErrors {
	errs := []ValidationError{}

	vValue := reflect.ValueOf(v)
	vType := reflect.TypeOf(v)

	if vValue.Kind() != reflect.Struct {
		errs = append(errs, ValidationError{
			Field: "Root",
			Err:   errors.New("Expected struct"),
		})
		return errs
	}

	for i := 0; i < vValue.NumField(); i++ {
		fieldValue := vValue.Field(i)
		fieldType := vType.Field(i)

		fieldErrs := validateField(fieldValue, fieldType)
		errs = append(errs, fieldErrs...)
	}

	return errs
}

// func main() {
// 	a := Validate()
// 	fmt.Println(a)
// }
