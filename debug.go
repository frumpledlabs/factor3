package factor3

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type fieldInfo struct {
	Key                 string
	EnvironmentVariable string
	DefaultValue        string
	CalculatedValue     interface{}
}

func debugFieldAndEnvironment(
	prefix string,
	input interface{},
) (map[string]fieldInfo, error) {
	fields := make(map[string]fieldInfo)

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	err := validateFieldInfo(inputType, inputValue)
	if err != nil {
		return fields, err
	}

	for i := 0; i < inputType.NumField(); i++ {

		field := inputValue.Field(i)
		fieldType := inputType.Field(i)
		name := inputType.Field(i).Name

		switch field.Kind() {
		case reflect.Ptr:
			println("Pointer field found o_o")
			debugFieldFromEnv(prefix, name, field.Elem(), fieldType)
		case reflect.Struct:
			reference := reflect.New(field.Type())

			structFields, err := debugFieldAndEnvironment(name, reference.Interface())
			for _, field := range structFields {
				fields[field.Key] = field
			}

			if err != nil {
				return fields, err
			}
		default:
			fieldInfo, err := debugFieldFromEnv(
				prefix,
				"",
				field,
				fieldType,
			)

			if err != nil {
				return fields, err
			}

			fields[fieldInfo.Key] = fieldInfo
		}
	}

	return fields, nil
}

func validateFieldInfo(inputType reflect.Type, input reflect.Value) error {
	var err error

	if inputType.Kind() != reflect.Struct {
		err = errors.New("Expected a struct")
		log.Error(
			"Error in debugEnvironmentInto()",
			map[string]interface{}{
				"msg": err.Error(),
			},
		)
	}

	return err
}

func validateField(fieldValue reflect.Value) error {
	var err error
	if !fieldValue.CanSet() {
		err = errors.New("field cannot be set")
	}

	return err
}

func debugFieldFromEnv(
	prefix string,
	envVar string,
	fieldValue reflect.Value,
	fieldType reflect.StructField,
) (fieldInfo, error) {

	var err error
	var defaultValue string
	var envVarOverride string

	var fieldInfo fieldInfo

	err = validateField(fieldValue)
	if err != nil {
		return fieldInfo, err
	}

	key := fmt.Sprintf("%s.%s", envVar, fieldType.Name)

	fieldInfo.Key = key

	envVar = fmt.Sprintf("%s_%s", envVar, fieldType.Name)
	envVar = macroCaser.Replace(envVar)

	tagDefinition, tagDefinitionExists := fieldType.Tag.Lookup(tagEnvName)
	var fieldData fieldData

	if tagDefinitionExists {
		fieldData = newFieldData(tagDefinition)
		if fieldData.keyIsOverriden {
			envVarOverride = fieldData.overrideKey
			fieldInfo.EnvironmentVariable = key
			println("Override key:", envVarOverride)
		}

		if fieldData.hasDefaultValue {
			defaultValue = fieldData.defaultValue
			fieldInfo.DefaultValue = defaultValue
		}
	}

	var envValue string
	envValue, err = debugEnvValueForField(fieldType, key)
	if err != nil {
		return fieldInfo, err
	}

	if isZeroValue(fieldValue) {
		if envValue == "" {
			envValue = defaultValue
		}

		if envValue == "" && fieldData.isRequired {
			return fieldInfo, errors.New("required field not set")
		}

		err = debugSetField(key, envValue, fieldValue)
		if err != nil {
			return fieldInfo, err
		}
	}

	return fieldInfo, nil
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

func debugSetField(key string, rawValue string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Struct:
		// return errors.New("Cannot set field value on type reflect.Struct")
	case reflect.Ptr:
		// return errors.New("Cannot set field value on type reflect.Ptr")
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
	}

	return nil
}
