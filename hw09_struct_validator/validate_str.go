package main

import (
	"errors"
	"reflect"
	"strconv"
)

const (
	len_cond = "len"
	regexp   = "regexp"
	in       = "in"
)

func validateLen(value string, condValue string) error {
	expectedLen, err := strconv.Atoi(condValue)
	if err != nil {
		return errors.New("Wrong tag format (len must be int)")
	}
	if valueLen := len(value); valueLen != expectedLen {
		return errors.New()
	}
}

func validateString(fieldValue reflect.Value, condName string, condValue string) error {
	var err error = nil
	switch condName {
	case len_cond:
		_ = 1
	case regexp:
		_ = 2
	case in:
		_ = 3
	}
}
