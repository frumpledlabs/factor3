package factor3

import (
	"errors"
	"fmt"
	"reflect"
)

// ReadEnvironmentInto environment into given configuration variable, using specific
// tags to determine requirements, values, and behavior.
func readEnvironmentInto(prefix string, input interface{}) error {
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return errors.New("Expected a struct pointer")
	}

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return errors.New("Expected a struct pointer")
	}

	for i := 0; i < inputType.NumField(); i++ {
		err := setFieldFromEnv(prefix, inputValue.Field(i), inputType.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func setFieldFromEnv(prefix string, field reflect.Value, fieldType reflect.StructField) error {
	var macroCaser = NewMacroCaseReplacer()

	if !field.CanSet() {
		log.Error("Field cannot be set.", "field", field) // TODO: Untested; determien how to get actual field name
		return errors.New("Field cannot be set.")
	}

	key := fmt.Sprintf("%s_%s", prefix, fieldType.Name)
	// originalKey := key
	key = macroCaser.Replace(key)

	envValue, err := getEnvValueForField(fieldType, key)
	if err != nil {
		return err
	}

	if isZeroValue(field) {
		err = setField(envValue, field)
		if err != nil {
			return err
		}
	}

	switch field.Kind() {
	case reflect.Ptr:
		setFieldFromEnv(key, field.Elem(), fieldType)
	case reflect.Struct:
		reference := reflect.New(field.Type())
		value := reference.Elem()

		value.Set(field)
		readEnvironmentInto(key, reference.Interface())
		field.Set(value)
	}

	log.Info("Set field value.", "field", fieldType.Name, "variable", key)

	return nil
}
