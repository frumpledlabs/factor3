package factor3

import (
	"errors"
	"os"
	"reflect"
	"strconv"
)

func isZeroValue(field reflect.Value) bool {
	return reflect.DeepEqual(
		reflect.Zero(field.Type()).Interface(),
		field.Interface(),
	)
}

func getEnvValueForField(field reflect.StructField, key string) (string, error) {
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

func setField(rawValue string, field reflect.Value) error {
	switch field.Kind() {
	case reflect.Bool:
		value, err := strconv.ParseBool(rawValue)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	case reflect.Float32:
		value, err := strconv.ParseFloat(rawValue, 32)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(float32(value)).Convert(field.Type()))
	case reflect.Float64:
		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	case reflect.Int:
		value, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int(value)).Convert(field.Type()))
	case reflect.Int8:
		value, err := strconv.ParseInt(rawValue, 10, 8)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int8(value)).Convert(field.Type()))
	case reflect.Int16:
		value, err := strconv.ParseInt(rawValue, 10, 16)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int16(value)).Convert(field.Type()))
	case reflect.Int32:
		value, err := strconv.ParseInt(rawValue, 10, 32)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int32(value)).Convert(field.Type()))
	case reflect.Int64:
		value, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	case reflect.String:
		field.Set(reflect.ValueOf(rawValue).Convert(field.Type()))
	case reflect.Struct:
		field.Set(reflect.New(field.Type()).Elem())
	case reflect.Ptr:
		field.Set(reflect.New(field.Type().Elem()))
	}

	return nil
}

func test(key string, field reflect.Value, fieldType reflect.StructField) {
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
}
