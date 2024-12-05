package hw09structvalidator

import (
	"reflect"
	"testing"

	"github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/mocks"
	"github.com/stretchr/testify/require"
)

func TestElementValidate(t *testing.T) {

	tasks := []struct {
		name       string
		fieldValue interface{}
		fieldType  reflect.Kind
	}{
		{name: "test_int", fieldValue: 10, fieldType: reflect.Int},
		// {name: "test_int8", fieldValue: 10, fieldType: reflect.Int8},
		// {name: "test_int16", fieldValue: 10, fieldType: reflect.Int16},
		// {name: "test_int32", fieldValue: 10, fieldType: reflect.Int32},
		// {name: "test_int64", fieldValue: 10, fieldType: reflect.Int64},
	}
	for _, tc := range tasks {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fieldValue := reflect.ValueOf(tc.fieldValue)

			condition := "min"
			conditionValue := "5"
			tag := condition + ":" + conditionValue

			intValidatorMock := mocks.NewIntValidatorInterface(t)
			intValidatorMock.EXPECT().Validate(fieldValue, condition, conditionValue).Return(nil)
			ev := ElementValidator{}
			ev.IntValidator = intValidatorMock

			err := ev.Validate(fieldValue, tc.fieldType, "testField", tag)
			require.NoError(t, err.Err)

		})
	}

	// ev := ElementValidator{}
	// // Тестирование для строкового типа
	// fieldValue := reflect.ValueOf("test")
	// fieldType := reflect.String
	// fieldName := "testField"
	// tag := "required"

	// err := ev.StringValidator.Validate(fieldValue, "required", "")
	// if err != nil {
	// 	t.Errorf("Ошибка при проверке строки: %v", err)
	// }

	// // Тестирование для числового типа
	// fieldValue = reflect.ValueOf(10)
	// fieldType = reflect.Int
	// fieldName = "testField"
	// tag = "min=5"

	// err = ev.IntValidator.Validate(fieldValue, "min", "5")
	// if err != nil {
	// 	t.Errorf("Ошибка при проверке числа: %v", err)
	// }

	// // Тестирование для неверного тега
	// fieldValue = reflect.ValueOf("test")
	// fieldType = reflect.String
	// fieldName = "testField"
	// tag = "invalid"

	// err = ev.StringValidator.Validate(fieldValue, "invalid", "")
	// if err == nil {
	// 	t.Errorf("Ожидается ошибка при проверке неверного тега")
	// }
}
