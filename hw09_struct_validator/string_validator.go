package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type StringValidator struct{}

func (strv StringValidator) validateLen(value string, condValue string) error {
	expectedLen, err := strconv.Atoi(condValue)
	if err != nil {
		return fmt.Errorf("wrong tag format: %s; input: %s", err, condValue)
	}
	valueLen := len(value)
	if valueLen != expectedLen {
		return fmt.Errorf("expected length: %d; real lenght: %d; input: %s", expectedLen, valueLen, value)
	}
	return nil
}

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

func (strv StringValidator) validateIn(value string, condValue string) error {
	for _, el := range strings.Split(value, ",") {
		el = strings.Trim(el, " ")
		if value == el {
			return nil
		}
	}
	return fmt.Errorf("inputed value (%s) is not in %s", value, condValue)
}

func (strv StringValidator) Validate(fieldValue reflect.Value, condName string, condValue string) error {
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
