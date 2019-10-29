package factor3

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const labelIsRequired = "required"

var envVariablePattern = regexp.MustCompile(`\${.*(?::-)?.*}`)

type fieldData struct {
	definition      string
	defaultValue    string
	overrideKey     string
	isRequired      bool
	keyIsOverriden  bool
	hasDefaultValue bool
}

func newFieldData(input string) fieldData {
	fd := fieldData{
		definition: input,
	}

	tags := strings.Split(input, ",")
	if len(tags) == 0 {
		return fd
	}

	fd.parseValuesDefinition()
	fd.parseKnownLabels()

	return fd
}

func (fd *fieldData) parseValuesDefinition() {
	values := strings.Split(fd.definition, ",")
	if len(values) == 0 {
		return
	}

	value := values[0]
	values = values[1:]

	if !envVariablePattern.MatchString(value) {
		return
	}

	envVarDefinition := newEnvVarDefinition(value)

	if envVarDefinition.varName != "" {
		fd.overrideKey = envVarDefinition.varName
		fd.keyIsOverriden = true
	}

	if envVarDefinition.defaultValue != "" {
		fd.defaultValue = envVarDefinition.defaultValue
		fd.hasDefaultValue = true
	}
}

func (fd *fieldData) parseKnownLabels() {
	for _, tag := range strings.Split(fd.definition, ",") {
		switch tag {
		case labelIsRequired:
			fd.isRequired = true
		}
	}
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
