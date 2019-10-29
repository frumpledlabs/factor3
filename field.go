package factor3

import (
	"errors"
	"os"
	"reflect"
	"strconv"
)

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(
		reflect.Zero(v.Type()).Interface(),
		v.Interface(),
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
	isRequired, _ := strconv.ParseBool(isRequiredValue)

	if isRequired && !isSet {
		return value, errors.New("No value set for required key: " + key)
	}

	return value, nil
}

func setField(rawValue string, v reflect.Value) error {
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
