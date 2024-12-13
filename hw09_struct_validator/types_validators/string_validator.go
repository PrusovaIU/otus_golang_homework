package types_validators

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
// error - ошибка, если длина строки не равна ожидаемой длине
func (strv StringValidator) validateLen(value string, condValue string) error {
	expectedLen, err := strconv.Atoi(strings.TrimSpace(condValue))
	if err != nil {
		return fmt.Errorf("wrong tag format: %s; input: %s", err, condValue)
	}
	valueLen := len(value)
	if valueLen != expectedLen {
		return fmt.Errorf("expected length: %d; real lenght: %d; input: %s", expectedLen, valueLen, value)
	}
	return nil
}

// validateRegexp проверяет, что строка соответствует регулярному выражению.
// Входные параметры:
// value - строка для проверки
// condValue - регулярное выражение
// Возвращаемое значение:
// error - ошибка, если строка не соответствует регулярному выражению
func (strv StringValidator) validateRegexp(value string, condValue string) error {
	condRegexp, err := regexp.Compile(condValue)
	if err != nil {
		return fmt.Errorf("invalid regexp: %s; input: %s", err, condValue)
	}
	if !condRegexp.Match([]byte(value)) {
		return fmt.Errorf("value does not matched to regexp; input: %s", value)
	}
	return nil
}

// validateIn проверяет, что значение находится в заданном списке.
// Входные параметры:
// value - значение для проверки
// condValue - список значений
// Возвращаемое значение:
// error - ошибка, если значение не находится в списке
func (strv StringValidator) validateIn(value string, condValue string) error {
	for _, el := range strings.Split(condValue, ",") {
		el = strings.TrimSpace(el)
		if value == el {
			return nil
		}
	}
	return fmt.Errorf("inputed value (%s) is not in %s", value, condValue)
}

// Validate проверяет значение типа string на соответствие заданному условию.
// Входные параметры:
// fieldValue - значение для проверки
// condName - имя условия
// condValue - значение условия
// Возвращаемое значение:
// error - ошибка, если значение не соответствует условию
func (strv StringValidator) Validate(fieldValue StringInterface, condName string, condValue string) error {
	var err error = nil
	value := fieldValue.String()
	switch condName {
	case len_cond:
		err = strv.validateLen(value, condValue)
	case regexp_cond:
		err = strv.validateRegexp(value, condValue)
	case in:
		err = strv.validateIn(value, condValue)
	}
	return err
}
