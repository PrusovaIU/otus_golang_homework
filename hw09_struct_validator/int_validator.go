package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type IntValidator struct{}

func (iv IntValidator) validateMinMax(value int64, condName string, condValue string) error {
	intCondValue, err := strconv.Atoi(condValue)
	if err != nil {
		return fmt.Errorf("wrong tag format: %s; input: %s", err, condValue)
	}
	switch condName {
	case min:
		if value < int64(intCondValue) {
			return fmt.Errorf("value must be less than %d; input: %d", intCondValue, value)
		}
	case max:
		if value > int64(intCondValue) {
			return fmt.Errorf("value must be greater than %d; input: %d", intCondValue, value)
		}
	}
	return nil
}

func (iv IntValidator) validateIn(value int64, condValue string) error {
	return nil
}

func (iv IntValidator) Validate(fieldValue reflect.Value, condName string, condValue string) error {
	var err error = nil
	value := fieldValue.Int()
	if condName == in {
		err = iv.validateIn(value, condValue)
	} else {
		err = iv.validateMinMax(value, condName, condValue)
	}
	return err
}
