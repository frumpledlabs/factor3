package factor3

import (
	"errors"
	"fmt"
	"reflect"
)

const tagEnvName = "env"

var macroCaser = newMacroCaseReplacer()

// readEnvironmentInto environment into given configuration variable, using specific
// tags to determine requirements, values, and behavior.
func readEnvironmentInto(prefix string, input interface{}) error {
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return errors.New("expected a pointer")
	}

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return errors.New("expected a struct")
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
	var err error
	var exists bool
	var key string
	var defaultValue string

	if !field.CanSet() {
		log.Error("Field cannot be set.",
			map[string]interface{}{
				"field": field,
			},
		)

		return errors.New("field cannot be set")
	}

	key = fmt.Sprintf("%s_%s", prefix, fieldType.Name)
	key = macroCaser.Replace(key)

	tagDefinition, exists := fieldType.Tag.Lookup(tagEnvName)
	var fieldData fieldData

	if exists {
		fieldData = newFieldData(tagDefinition)
		if fieldData.keyIsOverriden {
			key = fieldData.overrideKey
		}

		if fieldData.hasDefaultValue {
			defaultValue = fieldData.defaultValue
		}
	}

	var envValue string
	envValue, err = getEnvValueForField(fieldType, key)
	if err != nil {
		return err
	}

	if isZeroValue(field) {
		if envValue == "" {
			envValue = defaultValue
		}

		if envValue == "" && fieldData.isRequired {
			return errors.New("required field not set")
		}

		err = setField(envValue, field, defaultValue)
		if err != nil {
			return err
		}

		log.Info(
			"Set field value.",
			map[string]interface{}{
				"field":    fieldType.Name,
				"variable": key,
			},
		)
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

	log.Info(
		"Set field value.",
		map[string]interface{}{
			"field":    fieldType.Name,
			"variable": key,
		})

	return nil
}

// func setFieldValue(prefix string, field reflect.Value, fieldType reflect.StructField) error {

// }
