package factor3

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type fieldInfo struct {
	Key                 string
	EnvironmentVariable string
	DefaultValue        string
	CalculatedRawValue  interface{}
}

func debugReadEnvironmentInto(
	prefix string,
	input interface{},
) (map[string]fieldInfo, error) {
	fields := make(map[string]fieldInfo)

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	err := validateInput(inputType, inputValue)
	if err != nil {
		return fields, err
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputValue.Field(i)
		fieldType := inputType.Field(i)
		fieldName := inputType.Field(i).Name

		switch field.Kind() {
		case reflect.Struct:
			structFields, err := debugReadStruct(
				prefix,
				"."+fieldName,
				reflect.New(field.Type()).Interface(),
			)
			for _, field := range structFields {
				fields[field.Key] = field
			}

			if err != nil {
				return fields, err
			}
		default:
			fieldInfo, err := debugReadField(
				prefix,
				"",
				fieldName,
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

func debugReadStruct(
	envPrefix string,
	keyPrefix string,
	input interface{},
) (map[string]fieldInfo, error) {
	fields := make(map[string]fieldInfo)

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	err := validateInput(inputType, inputValue)
	if err != nil {
		return fields, err
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputValue.Field(i)
		fieldType := inputType.Field(i)
		fieldName := inputType.Field(i).Name

		switch field.Kind() {
		case reflect.Struct:

			structFields, err := debugReadStruct(
				envPrefix,
				keyPrefix,
				reflect.New(field.Type()).Interface(),
			)
			for _, field := range structFields {
				fields[field.Key] = field
			}

			if err != nil {
				return fields, err
			}
		default:
			fieldInfo, err := debugReadField(
				envPrefix,
				keyPrefix,
				fieldName,
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

func validateInput(fieldType reflect.Type, fieldValue reflect.Value) error {
	if fieldType.Kind() != reflect.Struct {
		return errors.New("Expected a struct")
	}

	return validateFieldCanBeSet(fieldValue)
}

func validateFieldCanBeSet(fieldValue reflect.Value) error {
	if !fieldValue.CanSet() {
		return errors.New("Field cannot be set")
	}

	return nil
}

func debugReadField(
	envPrefix string,
	keyPrefix string,
	name string,
	fieldValue reflect.Value,
	fieldType reflect.StructField,
) (fieldInfo, error) {
	var err error

	fieldInfo := fieldInfo{
		Key: fmt.Sprintf("%s.%s", keyPrefix, name),
	}

	err = validateFieldCanBeSet(fieldValue)
	if err != nil {
		return fieldInfo, err
	}

	envVar := fmt.Sprintf("%s_%s_%s", envPrefix, keyPrefix, name)
	envVar = macroCaser.Replace(envVar)

	tagDefinition, _ := fieldType.Tag.Lookup(tagEnvName)
	fieldData := newFieldData(tagDefinition)

	fieldInfo.EnvironmentVariable = envVar
	if fieldData.keyIsOverriden {
		fieldInfo.EnvironmentVariable = fieldData.overrideKey
	}

	if fieldData.hasDefaultValue {
		fieldInfo.DefaultValue = fieldData.defaultValue
	}

	return fieldInfo, nil
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
