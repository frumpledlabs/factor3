package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Loads configuration into the given variable, assuming all values are
// optional, unless otherwise tagged.
func Load(input interface{}) error {
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return errors.New("Expected a struct pointer")
	}

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return errors.New("Expected a struct pointer")
	}

	for i := 0; i < inputType.NumField(); i++ {
		err := setField("", inputValue.Field(i), inputType.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func setField(prefix string, field reflect.Value, fieldType reflect.StructField) error {
	if !field.CanSet() {
		return errors.New("Field cannot be set")
	}

	key := fmt.Sprintf("%s_%s", prefix, fieldType.Name)
	key = Replace(key)

	value := os.Getenv(key)
	isSet := len(value) > 0

	if !isSet {
		value = fieldType.Tag.Get("envDefault")
		isSet = len(value) > 0
	}

	isRequiredValue := fieldType.Tag.Get("envRequired")
	isRequired, err := strconv.ParseBool(isRequiredValue)
	if err != nil {
		// log.Warnf("Unrecognized tag '%s' for key: %s", isRequiredValue, key)
		isRequired = false
	}

	if isRequired && !isSet {
		return errors.New("No value set for required key: " + key)
	}

	if isZeroValue(field) {
		switch field.Kind() {
		case reflect.Bool:
			if val, err := strconv.ParseBool(value); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.Int:
			if val, err := strconv.ParseInt(value, 10, 64); err == nil {
				field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
			}
		case reflect.Int8:
			if val, err := strconv.ParseInt(value, 10, 8); err == nil {
				field.Set(reflect.ValueOf(int8(val)).Convert(field.Type()))
			}
		case reflect.Int16:
			if val, err := strconv.ParseInt(value, 10, 16); err == nil {
				field.Set(reflect.ValueOf(int16(val)).Convert(field.Type()))
			}
		case reflect.Int32:
			if val, err := strconv.ParseInt(value, 10, 32); err == nil {
				field.Set(reflect.ValueOf(int32(val)).Convert(field.Type()))
			}
		case reflect.Int64:
			if val, err := strconv.ParseInt(value, 10, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.Uint:
			if val, err := strconv.ParseUint(value, 10, 64); err == nil {
				field.Set(reflect.ValueOf(uint(val)).Convert(field.Type()))
			}
		case reflect.Uint8:
			if val, err := strconv.ParseUint(value, 10, 8); err == nil {
				field.Set(reflect.ValueOf(uint8(val)).Convert(field.Type()))
			}
		case reflect.Uint16:
			if val, err := strconv.ParseUint(value, 10, 16); err == nil {
				field.Set(reflect.ValueOf(uint16(val)).Convert(field.Type()))
			}
		case reflect.Uint32:
			if val, err := strconv.ParseUint(value, 10, 32); err == nil {
				field.Set(reflect.ValueOf(uint32(val)).Convert(field.Type()))
			}
		case reflect.Uint64:
			if val, err := strconv.ParseUint(value, 10, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.Uintptr:
			if val, err := strconv.ParseUint(value, 10, 64); err == nil {
				field.Set(reflect.ValueOf(uintptr(val)).Convert(field.Type()))
			}
		case reflect.Float32:
			if val, err := strconv.ParseFloat(value, 32); err == nil {
				field.Set(reflect.ValueOf(float32(val)).Convert(field.Type()))
			}
		case reflect.Float64:
			if val, err := strconv.ParseFloat(value, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.String:
			field.Set(reflect.ValueOf(value).Convert(field.Type()))
		case reflect.Struct:
			field.Set(reflect.New(field.Type()).Elem())
		case reflect.Ptr:
			field.Set(reflect.New(field.Type().Elem()))
		}
	}

	switch field.Kind() {
	case reflect.Ptr:
		setField(key, field.Elem(), fieldType)
	case reflect.Struct:
		reference := reflect.New(field.Type())
		value := reference.Elem()

		value.Set(field)
		Load(reference.Interface())
		field.Set(value)
	}

	return nil
}

func isZeroValue(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}

// Returns a slice of all field names' calculated env variable name
func Keys(configuration interface{}) ([]string, error) {
	return []string{}, errors.New("Keys() func is NOT yet implemented")
}
