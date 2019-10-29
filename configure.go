package factor3

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

const tagEnvName = "env"

var macroCaser = newMacroCaseReplacer()

// readEnvironmentInto environment into given configuration variable, using specific
// tags to determine requirements, values, and behavior.
func setFields(
	prefix string,
	input interface{},
) error {
	validateInput(input)

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	for i := 0; i < inputType.NumField(); i++ {
		err := setFieldFromEnv(prefix, inputValue.Field(i), inputType.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func validateInput(input interface{}) {
	inputType := reflect.ValueOf(input).Elem().Type()

	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		log.Fatal("Expected a pointer as input", nil)
	}

	if inputType.Kind() != reflect.Struct {
		log.Fatal("Expected a struct as input", nil)
	}
}

func validateField(key string, fieldValue reflect.Value) error {
	if !fieldValue.CanSet() {
		log.Error("Field cannot be set.",
			map[string]interface{}{
				"field": key,
			},
		)

		return errors.New("field cannot be set")
	}

	return nil
}

func setFieldFromEnv(
	prefix string,
	field reflect.Value,
	fieldType reflect.StructField,
) error {
	var err error
	var defaultValue string

	key := macroCaser.Replace(
		fmt.Sprintf("%s_%s", prefix, fieldType.Name),
	)

	err = validateField(key, field)
	if err != nil {
		return err
	}

	tagDefinition, _ := fieldType.Tag.Lookup(tagEnvName)

	fieldData := newFieldData(tagDefinition)
	if fieldData.keyIsOverriden {
		key = fieldData.overrideKey
	}

	if fieldData.hasDefaultValue {
		defaultValue = fieldData.defaultValue
	}

	var envValue string
	envValue, envValueIsSet := os.LookupEnv(key)

	if isZeroValue(field) {
		if !envValueIsSet {
			envValue = defaultValue
		}

		if envValue == "" && fieldData.isRequired {
			return errors.New("required field not set")
		}

		err = setFieldValue(envValue, field)
		if err != nil {
			return err
		}

		log.Debug(
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
		setFields(key, reference.Interface())
		field.Set(value)
	}

	return nil
}
