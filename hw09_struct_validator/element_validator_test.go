package hw09structvalidator

import (
	"errors"
	"reflect"
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/mocks"
	"github.com/stretchr/testify/require"
)

func TestElementValidate(t *testing.T) {
	int_tasks := []struct {
		name             string
		fieldValue       interface{}
		fieldType        reflect.Kind
		validationResult error
	}{
		{name: "test_int", fieldValue: 10, fieldType: reflect.Int, validationResult: nil},
		{name: "test_int8", fieldValue: 10, fieldType: reflect.Int8, validationResult: nil},
		{name: "test_int16", fieldValue: 10, fieldType: reflect.Int16, validationResult: nil},
		{name: "test_int32", fieldValue: 10, fieldType: reflect.Int32, validationResult: nil},
		{name: "test_int64", fieldValue: 10, fieldType: reflect.Int64, validationResult: nil},
		{name: "test_int_err", fieldValue: 10, fieldType: reflect.Int, validationResult: errors.New("test error")},
		{name: "test_int8_err", fieldValue: 10, fieldType: reflect.Int8, validationResult: errors.New("test error")},
		{name: "test_int16_err", fieldValue: 10, fieldType: reflect.Int16, validationResult: errors.New("test error")},
		{name: "test_int32_err", fieldValue: 10, fieldType: reflect.Int32, validationResult: errors.New("test error")},
		{name: "test_int64_err", fieldValue: 10, fieldType: reflect.Int64, validationResult: errors.New("test error")},
	}
	for _, tc := range int_tasks {
		t.Run(tc.name, func(t *testing.T) {
			fieldValue := reflect.ValueOf(tc.fieldValue)

			condition := "min"
			conditionValue := "5"
			tag := condition + ":" + conditionValue

			intValidatorMock := mocks.NewIntValidatorInterface(t)
			intValidatorMock.EXPECT().Validate(fieldValue, condition, conditionValue).Return(tc.validationResult)
			ev := ElementValidator{}
			ev.IntValidator = intValidatorMock

			err := ev.Validate(fieldValue, tc.fieldType, "testField", tag)
			if tc.validationResult == nil {
				require.NoError(t, err.Err)
			} else {
				require.Error(t, err.Err)
			}
		})
	}

	string_tasks := []struct {
		name             string
		validationResult error
	}{
		{name: "test_string", validationResult: nil},
		{name: "test_string_err", validationResult: errors.New("test error")},
	}
	for _, tc := range string_tasks {
		t.Run(tc.name, func(t *testing.T) {
			fieldValue := reflect.ValueOf("test")

			condition := "len"
			conditionValue := "4"
			tag := condition + ":" + conditionValue

			stringValidatorMock := mocks.NewStringValidatorInterface(t)
			stringValidatorMock.EXPECT().Validate(fieldValue, condition, conditionValue).Return(tc.validationResult)
			ev := ElementValidator{}
			ev.StringValidator = stringValidatorMock

			err := ev.Validate(fieldValue, reflect.String, "testField", tag)
			if tc.validationResult == nil {
				require.NoError(t, err.Err)
			} else {
				require.Error(t, err.Err)
			}
		})
	}

	t.Run("test_tag_error", func(t *testing.T) {
		fieldValue := reflect.ValueOf("test")
		tag := "wrong_format_tag"
		fieldName := "testField"

		ev := ElementValidator{}
		err := ev.Validate(fieldValue, reflect.String, fieldName, tag)
		require.Error(t, err.Err)
		require.Equal(t, err.Field, fieldName)
	})
}
