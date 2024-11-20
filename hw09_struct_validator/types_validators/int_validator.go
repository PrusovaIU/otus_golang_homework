package types_validators

import (
	"fmt"
	"strconv"
	"strings"
)

func checkGreat(value, minValue int64) error {
	if value < minValue {
		return fmt.Errorf("value must be great than %d; input: %d", minValue, value)
	}
	return nil
}

func checkLess(value, maxValue int64) error {
	if value > maxValue {
		return fmt.Errorf("value must be less than %d; input: %d", maxValue, value)
	}
	return nil
}

type IntInterface interface {
	Int() int64
}

type IntValidator struct{}

func (iv IntValidator) validateMinMax(value int64, condName string, condValue string) error {
	var checkFunc func(int64, int64) error = nil
	switch condName {
	case min:
		checkFunc = checkGreat
	case max:
		checkFunc = checkLess
	}
	if checkFunc != nil {
		intCondValue, err := strconv.Atoi(condValue)
		if err != nil {
			return fmt.Errorf("wrong tag format: %s; input: %s", err, condValue)
		}
		return checkFunc(value, int64(intCondValue))
	}
	return nil
}

func (iv IntValidator) validateIn(value int64, condValue string) error {
	split := strings.Split(condValue, ",")
	split_len := len(split)
	if split_len != 2 {
		return fmt.Errorf("condition must have only 2 values, but %d values have been get", split_len)
	}
	min_value, err := strconv.Atoi(strings.TrimSpace(split[0]))
	if err != nil {
		return fmt.Errorf("condition must be digit, but %s has been get", split[0])
	}
	max_value, err := strconv.Atoi(strings.TrimSpace(split[1]))
	if err != nil {
		return fmt.Errorf("condition must be digit, but %s has been get", split[0])
	}
	if value < int64(min_value) || value > int64(max_value) {
		return fmt.Errorf("value must be great than %d and less than %d; input: %d", max_value, max_value, value)
	}
	return nil
}

func (iv IntValidator) Validate(fieldValue IntInterface, condName string, condValue string) error {
	var err error = nil
	value := fieldValue.Int()
	if condName == in {
		err = iv.validateIn(value, condValue)
	} else {
		err = iv.validateMinMax(value, condName, condValue)
	}
	return err
}
