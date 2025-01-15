package types_validators

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
)

type StringInterface interface {
	String() string
}

type StringValidator struct{}

// validateLen проверяет, что длина строки равна ожидаемой длине.
// Входные параметры:
// value - строка для проверки
// condValue - ожидаемая длина строки
// Возвращаемое значение:
// error - ошибка, если длина строки не равна ожидаемой длине, или не удалось валидировать значение.
func (strv StringValidator) validateLen(value string, condValue string) error {
	expectedLen, err := strconv.Atoi(strings.TrimSpace(condValue))
	if err != nil {
		return fmt.Errorf("wrong tag format: %w; input: %s", err, condValue)
	}
	valueLen := len(value)
	if valueLen != expectedLen {
		// return fmt.Errorf("expected length: %d; real length: %d; input: %s", expectedLen, valueLen, value)
		return validationErrs.FieldValidationError{
			Message: fmt.Sprintf("expected length: %d; real length: %d; input: %s", expectedLen, valueLen, value),
		}
	}
	return nil
}

// validateRegexp проверяет, что строка соответствует регулярному выражению.
// Входные параметры:
// value - строка для проверки
// condValue - регулярное выражение
// Возвращаемое значение:
// error - ошибка, если строка не соответствует регулярному выражению, или не удалось валидировать значение.
func (strv StringValidator) validateRegexp(value string, condValue string) error {
	condRegexp, err := regexp.Compile(condValue)
	if err != nil {
		return fmt.Errorf("invalid regexp: %w; input: %s", err, condValue)
	}
	if !condRegexp.Match([]byte(value)) {
		// return fmt.Errorf("value does not matched to regexp; input: %s", value)
		return validationErrs.FieldValidationError{
			Message: fmt.Sprintf("value does not matched to regexp; input: %s", value),
		}
	}
	return nil
}

// validateIn проверяет, что значение находится в заданном списке.
// Входные параметры:
// value - значение для проверки
// condValue - список значений
// Возвращаемое значение:
// error - ошибка, если значение не находится в списке, или не удалось валидировать значение.
func (strv StringValidator) validateIn(value string, condValue string) error {
	for _, el := range strings.Split(condValue, ",") {
		el = strings.TrimSpace(el)
		if value == el {
			return nil
		}
	}
	// return fmt.Errorf("inputed value (%s) is not in %s", value, condValue)
	return validationErrs.FieldValidationError{
		Message: fmt.Sprintf("inputed value (%s) is not in %s", value, condValue),
	}
}

// Validate проверяет значение типа string на соответствие заданному условию.
// Входные параметры:
// fieldValue - значение для проверки
// condName - имя условия
// condValue - значение условия
// Возвращаемое значение:
// error - ошибка, если значение не соответствует условию, или не удалось валидировать значение.
func (strv StringValidator) Validate(fieldValue StringInterface, condName string, condValue string) error {
	var err error
	value := fieldValue.String()
	switch condName {
	case lenCond:
		err = strv.validateLen(value, condValue)
	case regexpCond:
		err = strv.validateRegexp(value, condValue)
	case in:
		err = strv.validateIn(value, condValue)
	}
	return err
}
