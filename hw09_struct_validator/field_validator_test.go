package hw09structvalidator

import (
	"errors"
	"reflect"
	"testing"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/mocks"
	"github.com/stretchr/testify/require"
)

func TestFieldValidatorElement(t *testing.T) {
	type TestElementStruct struct {
		TestValue string `validate:"len:4"`
	}

	tasks := []struct {
		name           string
		validateResult validationErrs.ValidationError
	}{
		{name: "test", validateResult: validationErrs.ValidationError{}},
		{
			name: "test_err", validateResult: validationErrs.ValidationError{
				Field: "testErr",
				Err:   errors.New("testErr"),
			},
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name+"element", func(t *testing.T) {
			testStruct := TestElementStruct{TestValue: "test"}
			fieldValue := reflect.ValueOf(testStruct).FieldByName("TestValue")
			fieldType, _ := reflect.TypeOf(testStruct).FieldByName("TestValue")

			elementValidatorMock := mocks.NewElementValidatorInterface(t)
			elementValidatorMock.
				EXPECT().
				Validate(fieldValue, fieldType.Type.Kind(), fieldType.Name, "len:4").
				Return(tc.validateResult)

			fv := FieldValidator{}
			fv.ElementValidator = elementValidatorMock

			errs := fv.Validate(fieldValue, fieldType)
			if tc.validateResult.IsErr() {
				require.Len(t, errs, 1)
			} else {
				require.Len(t, errs, 0)
			}
		})
	}
}

func TestFieldValidatorSlice(t *testing.T) {
	type TestSliceStrunct struct {
		TestValue []string `validate:"len:4"`
	}

	tasks := []struct {
		name           string
		validateResult validationErrs.ValidationErrors
	}{
		{name: "test", validateResult: []validationErrs.ValidationError{}},
		{name: "test_err", validateResult: []validationErrs.ValidationError{
			{
				Field: "testErr",
				Err:   errors.New("testErr"),
			},
		}},
	}

	for _, tc := range tasks {
		t.Run(tc.name+"slice", func(t *testing.T) {
			testStruct := TestSliceStrunct{TestValue: []string{"test"}}
			fieldValue := reflect.ValueOf(testStruct).FieldByName("TestValue")
			fieldType, _ := reflect.TypeOf(testStruct).FieldByName("TestValue")

			sliceValidatorMock := mocks.NewSliceValidatorInterface(t)
			sliceValidatorMock.EXPECT().Validate(fieldValue, fieldType, "len:4").Return(tc.validateResult)

			fv := FieldValidator{}
			fv.SliceValidator = sliceValidatorMock

			errs := fv.Validate(fieldValue, fieldType)
			if len(tc.validateResult) != 0 {
				require.Len(t, errs, 1)
			} else {
				require.Len(t, errs, 0)
			}
		})
	}
}
