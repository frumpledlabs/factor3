package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Test struct {
	Defaults string `envDefault:"default_value"`
	Nested   struct {
		Value string
	}
	Deep struct {
		Nested struct {
			Value string `env:"required"`
		}
	}
}

func main() {

	var conf Test

	printFields(&conf)
	printFields(&conf.Nested)
	printFields(&conf.Deep.Nested)

	//Load("", conf)
	fmt.Println("Conf:")
	fmt.Println(conf)
}

func printFields(input interface{}) {
	var s reflect.Value
	s = reflect.ValueOf(input)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	fmt.Println("\t\tNumFields:", s.NumField())

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		field := s.Type().Field(i)

		//fmt.Println(s.Type().Field(i).Name)
		//fmt.Println("\tType:\t\t\t", f.Type())
		//fmt.Println("\tCanAddr:\t\t", reflect.ValueOf(f).CanAddr())
		//fmt.Println("\tInterface:\t\t", f.Interface())
		//fmt.Println("\tIsValid:\t\t", f.IsValid())
		//fmt.Println("\tKind:   \t\t", f.Type())

		if f.Type().Kind() == reflect.String {
			fmt.Println("Field:", field.Name, "- set")
			f.SetString("PASS")
		} else if f.Type().Kind() == reflect.Struct {
			fmt.Println("Field:", field.Name, "- Iterating on:", reflect.ValueOf(field).Kind())
			printFields(f)
		} else {
			fmt.Println("\tSkipping field (", field.Name, ") of unknown type:", f.Type().Kind())
		}
	}
}

func Load(prefix string, input interface{}) {
	parseValue(prefix, input)
}

func parseValue(prefix string, input interface{}) {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
		inputValue = inputValue.Elem()
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			parseValue(prefix+"_"+field.Name, fieldValue.Interface())
		case reflect.String:
			name := strings.Trim(fmt.Sprintf(prefix+"_"+field.Name), "_")
			fmt.Println("Field Index:", name)
			setField(field)
			setValue(inputValue, field)
		default:
			// TODO: log warn: skipping unexpected value of type: fieldValue.Kind()
			fmt.Println("UNKNOWN Kind:", fieldValue.Kind())
		}
	}
}

func setField(field reflect.StructField) {
	fmt.Println("Setting field", field.Name)
	fieldValue := reflect.ValueOf(field)

	if !fieldValue.CanSet() {
		fmt.Println("\tCANNOT set")
		return
	}

	fmt.Println("\tCAN set field:", field)

	fmt.Println("Type:", reflect.TypeOf(field))
	fmt.Println("Kind:", reflect.TypeOf(field).Kind())
	fmt.Println("Field:", field.Name)
	fmt.Println("Tags:", field.Tag.Get("env"), field.Tag.Get("envDefault"))

	return
}

func setValue(value reflect.Value, field reflect.StructField) {
	fmt.Println("Setting value", field.Name)
	fmt.Println("\tIsValid:", value.IsValid())
	fmt.Println("\tKind:", value.Type().Kind())
	if !value.IsValid() || value.Type().Kind() != reflect.String {
		fmt.Println("\tCANNOT set:", field.Name)
		return
	}
	fmt.Println("\tCAN set:", field.Name)
	value.SetString("PASS")
}
