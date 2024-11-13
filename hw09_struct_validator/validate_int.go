package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Int interface {
	int8 | int16 | int32 | int64
}

func validateMinMax[T Int](value T, condName string, condValue string) error {
	intCondValue, err := strconv.Atoi(condValue)
	if err != nil {
		return fmt.Errorf("wrong tag format: %s; input: %s", err, condValue)
	}
	switch condName {
	case min:
		if value < T(intCondValue) {
			return fmt.Errorf("value must be less than %d; input: %d", intCondValue, value)
		}
	case max:
		if value > T(intCondValue) {
			return fmt.Errorf("value must be greater than %d; input: %d", intCondValue, value)
		}
	}
	return nil
}

func validateIntIn[T Int](value T, condValue string) error {
	return nil
}

func validateInt(fieldValue reflect.Value, condName string, condValue string) error {
	var err error = nil
	value := fieldValue.Int()
	if condName == in {
		err = validateIntIn(value, condValue)
	} else {
		err = validateMinMax(value, condName, condValue)
	}
	return err
}
