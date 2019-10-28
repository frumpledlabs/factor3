package factor3

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// TODO:  Replace environment file w/ this
// TODO:  Rename this to not have debug in name
// TODO:  Add debug flag to determine whether to set value or not
// TODO:  Add debug logging (flaggable on/off feature) for all fields found w/ this app

type fieldInfo struct {
	FieldValue          reflect.Value
	Key                 string      `json:"key"`
	EnvironmentVariable string      `json:"environment_variable"`
	DefaultValue        string      `json:"default_value"`
	CalculatedRawValue  interface{} `json:"calculated_raw_value"`
}

func loadFieldsFromEnvironmentFor(
	prefix string,
	input interface{},
) (map[string]fieldInfo, error) {
	fields := make(map[string]fieldInfo)

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return fields, errors.New("Expected a struct")
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
			fieldInfo, err := readField(
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
		fieldValue := inputValue.Field(i)
		fieldType := inputType.Field(i)
		fieldName := inputType.Field(i).Name

		switch fieldValue.Kind() {
		case reflect.Struct:
			keyPrefix = fmt.Sprintf(
				"%s.%s",
				keyPrefix,
				fieldName,
			)

			structFields, err := debugReadStruct(
				envPrefix,
				keyPrefix,
				reflect.New(fieldValue.Type()).Interface(),
			)
			for _, field := range structFields {
				fields[field.Key] = field
			}

			if err != nil {
				return fields, err
			}
		default:
			fieldInfo, err := readField(
				envPrefix,
				keyPrefix,
				fieldName,
				fieldValue,
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

func readField(
	envPrefix string,
	keyPrefix string,
	name string,
	fieldValue reflect.Value,
	fieldType reflect.StructField,
) (fieldInfo, error) {
	var err error

	fieldInfo := fieldInfo{
		FieldValue: fieldValue,
		Key:        fmt.Sprintf("%s.%s", keyPrefix, name),
	}

	err = validateFieldCanBeSet(fieldValue)
	if err != nil {
		return fieldInfo, err
	}

	envVar := macroCaser.Replace(
		fmt.Sprintf("%s_%s_%s", envPrefix, keyPrefix, name),
	)

	tagDefinition, _ := fieldType.Tag.Lookup(tagEnvName)
	fieldData := newFieldData(tagDefinition)

	if fieldData.hasDefaultValue {
		fieldInfo.DefaultValue = fieldData.defaultValue
		fieldInfo.CalculatedRawValue = fieldData.defaultValue
	}

	fieldInfo.EnvironmentVariable = envVar
	if fieldData.keyIsOverriden {
		fieldInfo.EnvironmentVariable = fieldData.overrideKey
		value, isSet := os.LookupEnv(fieldInfo.EnvironmentVariable)
		if !isSet {
			value = fieldInfo.DefaultValue
		}
		fieldInfo.CalculatedRawValue = value
	}

	return fieldInfo, nil
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
