package factor3

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
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

	switch field.Kind() {
	case reflect.Struct:
		reference := reflect.New(field.Type())
		value := reference.Elem()

		value.Set(field)
		setFields(key, reference.Interface())
		field.Set(value)

	default:
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
				return errors.New("Required field not set")
			}

			err = setFieldValue(envValue, field)
			if err != nil {
				return err
			}

			log.Debug(
				"Set field value.",
				map[string]interface{}{
					"field": fieldType.Name,
					"key":   key,
				},
			)
		}
	}

	return nil
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(
		reflect.Zero(v.Type()).Interface(),
		v.Interface(),
	)
}

func setFieldValue(rawValue string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Bool:
		value, err := strconv.ParseBool(rawValue)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(value).Convert(v.Type()))
	case reflect.Float32:
		value, err := strconv.ParseFloat(rawValue, 32)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(float32(value)).Convert(v.Type()))
	case reflect.Float64:
		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(value).Convert(v.Type()))
	case reflect.Int:
		value, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(int(value)).Convert(v.Type()))
	case reflect.Int8:
		value, err := strconv.ParseInt(rawValue, 10, 8)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(int8(value)).Convert(v.Type()))
	case reflect.Int16:
		value, err := strconv.ParseInt(rawValue, 10, 16)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(int16(value)).Convert(v.Type()))
	case reflect.Int32:
		value, err := strconv.ParseInt(rawValue, 10, 32)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(int32(value)).Convert(v.Type()))
	case reflect.Int64:
		value, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(value).Convert(v.Type()))
	case reflect.String:
		v.Set(reflect.ValueOf(rawValue).Convert(v.Type()))
	case reflect.Struct:
		v.Set(reflect.New(v.Type()).Elem())
	case reflect.Ptr:
		v.Set(reflect.New(v.Type().Elem()))
	}

	return nil
}
