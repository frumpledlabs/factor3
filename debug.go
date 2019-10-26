package factor3

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func debugEnvironmentInto(
	prefix string,
	input interface{},
) (map[string]interface{}, error) {
	println("Debug env")

	output := make(map[string]interface{})

	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return output, errors.New("expected a pointer")
	}

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return output, errors.New("expected a struct")
	}

	println("debug fields")
	for i := 0; i < inputType.NumField(); i++ {
		err := debugFieldFromEnv(prefix, inputValue.Field(i), inputType.Field(i))
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func debugFieldFromEnv(prefix string, field reflect.Value, fieldType reflect.StructField) error {
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
	envValue, err = debugEnvValueForField(fieldType, key)
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

		err = debugField(envValue, field)
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
		debugFieldFromEnv(key, field.Elem(), fieldType)
	case reflect.Struct:
		reference := reflect.New(field.Type())
		value := reference.Elem()

		value.Set(field)
		debugEnvironmentInto(key, reference.Interface())
		field.Set(value)
	}

	return nil
}

func debugEnvValueForField(field reflect.StructField, key string) (string, error) {
	value := os.Getenv(key)
	isSet := len(value) > 0

	if !isSet {
		value = field.Tag.Get("envDefault")
		isSet = len(value) > 0
	}

	isRequiredValue := field.Tag.Get("envRequired")
	isRequired, err := strconv.ParseBool(isRequiredValue)
	if err != nil {
		// log.Warnf("Unrecognized tag '%s' for key: %s", isRequiredValue, key)
		isRequired = false
	}

	if isRequired && !isSet {
		return value, errors.New("No value set for required key: " + key)
	}

	return value, nil
}

func debugField(rawValue string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Bool:
		value, err := strconv.ParseBool(rawValue)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(value).Convert(v.Type()))
		println("debug value:", value)
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
