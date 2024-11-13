package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func validateLen(value string, condValue string) error {
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

func validateRegexp(value string, condValue string) error {
	condRegexp, err := regexp.Compile(condValue)
	if err != nil {
		return fmt.Errorf("invalid regexp: %s; input: %s", err, condValue)
	}
	if !condRegexp.Match([]byte(value)) {
		return fmt.Errorf("value does not matched to regexp; input: %s", value)
	}
	return nil
}

func validateIn(value string, condValue string) error {
	for _, el := range strings.Split(value, ",") {
		el = strings.Trim(el, " ")
		if value == el {
			return nil
		}
	}
	return fmt.Errorf("inputed value (%s) is not in %s", value, condValue)
}

func validateString(fieldValue reflect.Value, condName string, condValue string) error {
	var err error = nil
	value := fieldValue.String()
	switch condName {
	case len_cond:
		err = validateLen(value, condValue)
	case regexp_cond:
		err = validateRegexp(value, condValue)
	case in:
		err = validateIn(value, condValue)
	}
	return err
}
