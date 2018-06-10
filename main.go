package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Test struct {
	Root   string `envDefault:"default_value"`
	Nested struct {
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

	Load("", &conf)
	//Load("", &conf.Nested)
	//Load("", &conf.Deep.Nested)

	fmt.Println("Conf:")
	fmt.Println(conf)
}

func Load(prefix string, input interface{}) {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
		inputValue = inputValue.Elem()
		fmt.Println("\tchanged to reference")
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			Load(prefix+"_"+field.Name, fieldValue.Interface())
		case reflect.String:
			name := strings.Trim(fmt.Sprintf(prefix+"_"+field.Name), "_")
			fmt.Println(name)
			//fmt.Println(field.Tag.Get("env"))
			//fmt.Println(field.Tag.Get("envDefault"))
			if fieldValue.IsValid() {
				if fieldValue.CanSet() {
					fieldValue.SetString("PASS")
				}
			}
		default:
			// TODO: log warn: skipping unexpected value of type: fieldValue.Kind()
			fmt.Println("UNKNOWN Kind:", fieldValue.Kind())
		}
	}
}
