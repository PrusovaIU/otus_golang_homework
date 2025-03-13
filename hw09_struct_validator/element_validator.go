package hw09structvalidator

import (
	"errors"
	"reflect"
	"strings"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/types_validators"
)

type IntValidatorInterface interface {
	Validate(types_validators.IntInterface, string, string) error
}

type StringValidatorInterface interface {
	Validate(types_validators.StringInterface, string, string) error
}

type ElementValidator struct {
	IntValidator    IntValidatorInterface
	StringValidator StringValidatorInterface
}

func NewElementValidator() ElementValidator {
	ev := ElementValidator{}
	ev.IntValidator = types_validators.IntValidator{}
	ev.StringValidator = types_validators.StringValidator{}
	return ev
}

// parseTag разбирает тег валидации на условие и значение.
// Входной параметр:
// tag - тег валидации
// Возвращаемые значения:
// condition - условие валидации
// value - значение для проверки
// error - ошибка, если тег валидации имеет неверный формат.
func (ev ElementValidator) parseTag(tag string) (string, string, error) {
	split := strings.Split(tag, ":")
	if len(split) != 2 {
		return "", "", errors.New("invalid tag format")
	}
	condition := strings.TrimSpace(split[0])
	value := strings.TrimSpace(split[1])
	return condition, value, nil
}

// Validate проверяет значение поля на соответствие заданному в теге условию.
// Входные параметры:
// fieldValue - значение поля
// fieldType - тип поля
// fieldName - имя поля
// tag - тег валидации
// Возвращаемое значение:
// ValidationError - структура с информацией о возможной ошибке валидации.
func (ev ElementValidator) Validate(
	fieldValue reflect.Value,
	fieldType reflect.Kind,
	fieldName string,
	tag string,
) error {
	var err error
	condition, condition_value, err := ev.parseTag(tag)
	if err == nil {
		switch fieldType {
		case reflect.String:
			err = ev.StringValidator.Validate(fieldValue, condition, condition_value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			err = ev.IntValidator.Validate(fieldValue, condition, condition_value)
		}
	}
	if elValidatorErr, ok := err.(validationErrs.FieldValidationError); ok {
		return validationErrs.ValidationError{
			Field: fieldName,
			Err:   elValidatorErr,
		}
	}
	return err
}
