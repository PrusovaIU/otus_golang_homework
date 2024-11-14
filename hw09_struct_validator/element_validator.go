package main

import (
	"errors"
	"reflect"
	"strings"
)

type TypeValidatorInterface interface {
	Validate(reflect.Value, string, string) error
}

type ElementValidator struct {
	IntValidator    TypeValidatorInterface
	StringValidator TypeValidatorInterface
}

func NewElementValidator() ElementValidator {
	ev := ElementValidator{}
	ev.IntValidator = IntValidator{}
	ev.StringValidator = StringValidator{}
	return ev
}

func (ev ElementValidator) parseTag(tag string) (string, string, error) {
	split := strings.Split(tag, ":")
	if len(split) != 2 {
		return "", "", errors.New("wrong tag format")
	}
	condition := strings.Trim(split[0], " ")
	value := strings.Trim(split[1], " ")
	return condition, value, nil
}

func (ev ElementValidator) Validate(fieldValue reflect.Value, fieldType reflect.Kind, fieldName string, tag string) ValidationError {
	var err error = nil
	var validationErr = ValidationError{}
	condition, condition_value, err := ev.parseTag(tag)
	if err == nil {
		switch fieldType {
		case reflect.String:
			err = ev.StringValidator.Validate(fieldValue, condition, condition_value)
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			err = ev.IntValidator.Validate(fieldValue, condition, condition_value)
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
