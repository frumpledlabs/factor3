package factor3

import (
	"errors"
	"fmt"
	"reflect"
)

// Loads environment into given configuration variable, using specific
// tags to determine requirements, values, and behavior.
func Load(input interface{}) error {
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return errors.New("Expected a struct pointer")
	}

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return errors.New("Expected a struct pointer")
	}

	for i := 0; i < inputType.NumField(); i++ {
		err := setFieldFromEnv("", inputValue.Field(i), inputType.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func setFieldFromEnv(prefix string, field reflect.Value, fieldType reflect.StructField) error {
	var macroCaser = NewMacroCaseReplacer()

	if !field.CanSet() {
		return errors.New("Field cannot be set")
	}

	key := fmt.Sprintf("%s_%s", prefix, fieldType.Name)
	key = macroCaser.Replace(key)

	envValue, err := getEnvValueForField(fieldType, key)
	if err != nil {
		return err
	}

	if isZeroValue(field) {
		setField(envValue, field)
	}

	switch field.Kind() {
	case reflect.Ptr:
		setFieldFromEnv(key, field.Elem(), fieldType)
	case reflect.Struct:
		reference := reflect.New(field.Type())
		value := reference.Elem()

		value.Set(field)
		Load(reference.Interface())
		field.Set(value)
	}

	return nil
}
