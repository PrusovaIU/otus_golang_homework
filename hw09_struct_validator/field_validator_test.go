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
		validateResult error
	}{
		{name: "test", validateResult: nil},
		{
			name: "test_validation_err", validateResult: validationErrs.ValidationError{
				Field: "testErr",
				Err:   errors.New("testErr"),
			},
		},
	}

	testStruct := TestElementStruct{TestValue: "test"}
	fieldValue := reflect.ValueOf(testStruct).FieldByName("TestValue")
	fieldType, _ := reflect.TypeOf(testStruct).FieldByName("TestValue")

	for _, tc := range tasks {
		t.Run(tc.name+"element", func(t *testing.T) {
			elementValidatorMock := mocks.NewElementValidatorInterface(t)
			elementValidatorMock.
				EXPECT().
				Validate(fieldValue, fieldType.Type.Kind(), fieldType.Name, "len:4").
				Return(tc.validateResult)

			fv := FieldValidator{}
			fv.ElementValidator = elementValidatorMock

			validationErrs, err := fv.Validate(fieldValue, fieldType)
			require.NoError(t, err)
			if tc.validateResult != nil {
				require.Len(t, validationErrs, 1)
			} else {
				require.Len(t, validationErrs, 0)
			}
		})
	}

	t.Run("test_error", func(t *testing.T) {
		elementValidatorMock := mocks.NewElementValidatorInterface(t)
		elementValidatorMock.
			EXPECT().
			Validate(fieldValue, fieldType.Type.Kind(), fieldType.Name, "len:4").
			Return(errors.New("Test error"))

		fv := FieldValidator{}
		fv.ElementValidator = elementValidatorMock

		validationErrs, err := fv.Validate(fieldValue, fieldType)
		require.Error(t, err)
		require.Nil(t, validationErrs)
	})
}

func TestFieldValidatorSlice(t *testing.T) {
	type TestSliceStrunct struct {
		TestValue []string `validate:"len:4"`
	}

	tasks := []struct {
		name           string
		validateResult validationErrs.ValidationErrors
	}{
		{name: "test_", validateResult: []validationErrs.ValidationError{}},
		{name: "test_validation_errs_", validateResult: []validationErrs.ValidationError{
			{
				Field: "testErr",
				Err:   errors.New("testErr"),
			},
		}},
	}

	testStruct := TestSliceStrunct{TestValue: []string{"test"}}
	fieldValue := reflect.ValueOf(testStruct).FieldByName("TestValue")
	fieldType, _ := reflect.TypeOf(testStruct).FieldByName("TestValue")

	for _, tc := range tasks {
		t.Run(tc.name+"slice", func(t *testing.T) {
			sliceValidatorMock := mocks.NewSliceValidatorInterface(t)
			sliceValidatorMock.EXPECT().Validate(fieldValue, fieldType, "len:4").Return(tc.validateResult, nil)

			fv := FieldValidator{}
			fv.SliceValidator = sliceValidatorMock

			validationErrs, err := fv.Validate(fieldValue, fieldType)
			require.NoError(t, err)
			require.Len(t, validationErrs, len(tc.validateResult))
		})
	}

	t.Run("test_error", func(t *testing.T) {
		sliceValidatorMock := mocks.NewSliceValidatorInterface(t)
		sliceValidatorMock.EXPECT().Validate(fieldValue, fieldType, "len:4").Return(nil, errors.New("Test error"))

		fv := FieldValidator{}
		fv.SliceValidator = sliceValidatorMock

		validationErrs, err := fv.Validate(fieldValue, fieldType)
		require.Error(t, err)
		require.Nil(t, validationErrs)
	})
}
