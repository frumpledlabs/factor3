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

func debugEnvironmentInto(
	prefix string,
	input interface{},
) (map[string]fieldInfo, error) {
	output := make(map[string]fieldInfo)

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	// err := validateFieldInfo(inputType, inputValue)
	// if err != nil {
	// 	return output, err
	// }

	for i := 0; i < inputType.NumField(); i++ {
		name := inputType.Field(i).Name
		println("prefix:", prefix, "-", name)
		fields, err := debugFieldFromEnv(prefix, "", inputValue.Field(i), inputType.Field(i))
		if err != nil {
			return output, err
		}
		for key, fieldInfo := range fields {
			output[key] = fieldInfo
		}
	}

	return output, nil
}

// func validateFieldInfo(inputType reflect.Type, input reflect.Value) error {
// 	var err error
// 	if (reflect.TypeOf(input).Kind() != reflect.Ptr) ||
// 		(inputType.Kind() != reflect.Struct) {
// 		err = errors.New("Expected a struct pointer")
// 		log.Error(
// 			"Error in debugEnvironmentInto()",
// 			map[string]interface{}{
// 				"msg": err.Error(),
// 			},
// 		)
// 	}

// 	return err
// }

func debugFieldFromEnv(
	prefix string,
	envVar string,
	field reflect.Value,
	fieldType reflect.StructField,
) (map[string]fieldInfo, error) {

	var err error
	var defaultValue string
	var envVarOverride string

	fields := make(map[string]fieldInfo)
	var fieldInfo fieldInfo

	if !field.CanSet() {
		log.Error("Field cannot be set.",
			map[string]interface{}{
				"field": field,
			},
		)

		return fields, errors.New("field cannot be set")
	}

	key := fmt.Sprintf("%s.%s", envVar, fieldType.Name)
	println("key:", key)

	switch field.Kind() {
	case reflect.Ptr:
		println("Pointer field found o_o")
		debugFieldFromEnv(prefix, key, field.Elem(), fieldType)
	case reflect.Struct:
		reference := reflect.New(field.Type())

		fields, err := debugEnvironmentInto(key, reference.Interface())
		for key, field := range fields {
			fields[key] = field
		}

		return fields, err
	}

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
			println("key:", key)
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
		return fields, err
	}

	println("key:", key)

	if isZeroValue(field) {
		if envValue == "" {
			envValue = defaultValue
		}

		if envValue == "" && fieldData.isRequired {
			return fields, errors.New("required field not set")
		}

		err = debugField(key, envValue, field)
		if err != nil {
			return fields, err
		}

		log.Debug(
			"Set field value.",
			map[string]interface{}{
				"field":    fieldType.Name,
				"variable": key,
			},
		)
	}

	fields[key] = fieldInfo

	for key := range fields {
		println("KEY:", key)
	}

	return fields, nil
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

func debugField(key string, rawValue string, v reflect.Value) error {

	if v.Kind() != reflect.Struct && (v.Kind() != reflect.Ptr) {
		println(key, ":", rawValue)
	}

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
