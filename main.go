package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Load(prefix string, input interface{}) error {
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return errors.New("Expected a struct pointer")
	}

	inputValue := reflect.ValueOf(input).Elem()
	inputType := inputValue.Type()

	if inputType.Kind() != reflect.Struct {
		return errors.New("Expected a struct pointer")
	}

	for i := 0; i < inputType.NumField(); i++ {
		defaultValue := inputType.Field(i).Tag.Get("envDefault")
		setField(prefix, inputValue.Field(i), inputType.Field(i), defaultValue)
	}

	return nil
}

func setField(prefix string, field reflect.Value, fieldType reflect.StructField, defaultVal string) error {
	if !field.CanSet() {
		return errors.New("Field cannot be set")
	}

	prefix = fmt.Sprintf("%s_%s", prefix, fieldType.Name)
	fmt.Print("PREFIX:", prefix)
	prefix = Replace(prefix)
	fmt.Println("->", prefix)

	fmt.Println("Default Value:", defaultVal)
	fmt.Println("Prefix:", prefix)
	fmt.Println("Checking ENV for value:", prefix)
	v := os.Getenv(prefix)
	fmt.Println("Found ENV:", v)
	isSet := len(v) > 0

	isRequired := fieldType.Tag.Get("env") == "required"

	if !isSet && isRequired {
		return errors.New("Missing value for key: " + prefix)
	}

	defaultVal = v
	fmt.Println("Set value:", defaultVal)

	//if isZeroValue(field) {
	//	fmt.Println(prefix)
	//	fmt.Printf("\tPre-field:\t%+v\n", field)
	//}

	if isZeroValue(field) {
		switch field.Kind() {
		case reflect.Bool:
			if val, err := strconv.ParseBool(defaultVal); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.Int:
			if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
				field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
			}
		case reflect.Int8:
			if val, err := strconv.ParseInt(defaultVal, 10, 8); err == nil {
				field.Set(reflect.ValueOf(int8(val)).Convert(field.Type()))
			}
		case reflect.Int16:
			if val, err := strconv.ParseInt(defaultVal, 10, 16); err == nil {
				field.Set(reflect.ValueOf(int16(val)).Convert(field.Type()))
			}
		case reflect.Int32:
			if val, err := strconv.ParseInt(defaultVal, 10, 32); err == nil {
				field.Set(reflect.ValueOf(int32(val)).Convert(field.Type()))
			}
		case reflect.Int64:
			if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.Uint:
			if val, err := strconv.ParseUint(defaultVal, 10, 64); err == nil {
				field.Set(reflect.ValueOf(uint(val)).Convert(field.Type()))
			}
		case reflect.Uint8:
			if val, err := strconv.ParseUint(defaultVal, 10, 8); err == nil {
				field.Set(reflect.ValueOf(uint8(val)).Convert(field.Type()))
			}
		case reflect.Uint16:
			if val, err := strconv.ParseUint(defaultVal, 10, 16); err == nil {
				field.Set(reflect.ValueOf(uint16(val)).Convert(field.Type()))
			}
		case reflect.Uint32:
			if val, err := strconv.ParseUint(defaultVal, 10, 32); err == nil {
				field.Set(reflect.ValueOf(uint32(val)).Convert(field.Type()))
			}
		case reflect.Uint64:
			if val, err := strconv.ParseUint(defaultVal, 10, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.Uintptr:
			if val, err := strconv.ParseUint(defaultVal, 10, 64); err == nil {
				field.Set(reflect.ValueOf(uintptr(val)).Convert(field.Type()))
			}
		case reflect.Float32:
			if val, err := strconv.ParseFloat(defaultVal, 32); err == nil {
				field.Set(reflect.ValueOf(float32(val)).Convert(field.Type()))
			}
		case reflect.Float64:
			if val, err := strconv.ParseFloat(defaultVal, 64); err == nil {
				field.Set(reflect.ValueOf(val).Convert(field.Type()))
			}
		case reflect.String:
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		case reflect.Struct:
			field.Set(reflect.New(field.Type()).Elem())
		case reflect.Ptr:
			field.Set(reflect.New(field.Type().Elem()))
		}
	}

	switch field.Kind() {
	case reflect.Ptr:
		setField(prefix, field.Elem(), fieldType, defaultVal)
	case reflect.Struct:
		reference := reflect.New(field.Type())
		value := reference.Elem()

		value.Set(field)
		Load(prefix, reference.Interface())
		field.Set(value)
	}
	//fmt.Printf("\tPost-field:\t%+v\n", field)

	return nil
}

func isZeroValue(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}
