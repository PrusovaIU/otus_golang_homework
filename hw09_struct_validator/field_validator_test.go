package hw09structvalidator

import (
	"errors"
	"reflect"
	"testing"

	validationErrs "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"
	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/mocks"
	"github.com/stretchr/testify/require"
)

func TestFieldValidator(t *testing.T) {
	elementTasks := []struct {
		name           string
		validateResult validationErrs.ValidationError
	}{
		{name: "test", validateResult: validationErrs.ValidationError{}},
		{name: "test", validateResult: validationErrs.ValidationError{
			Field: "testErr",
			Err:   errors.New("testErr"),
		},
		},
	}
	for _, tc := range elementTasks {
		t.Run(tc.name, func(t *testing.T) {
			type TestElementStruct struct {
				testValue string `validate:"len:4"`
			}
			testStruct := TestElementStruct{testValue: "test"}
			fieldValue := reflect.ValueOf(testStruct).FieldByName("testValue")
			fieldType, _ := reflect.TypeOf(testStruct).FieldByName("testValue")

			elementValidatorMock := mocks.NewElementValidatorInterface(t)
			elementValidatorMock.EXPECT().Validate(fieldValue, fieldType.Type.Kind(), fieldType.Name, "len:4").Return(tc.validateResult)

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
