package main

import (
	"fmt"
	"reflect"
)

func main() {
	type test struct {
		Defaults string `envDefault:"default_value"`
		Nested   struct {
			Deep struct {
				Value string `env:"required"`
			}
		}
	}

	var conf test
	LoadEnvironment("", conf)

	fmt.Println("Conf:")
	fmt.Println(conf)
}

func LoadEnvironment(prefix string, input interface{}) {
	parseValue(prefix, input)
}

func parseValue(prefix string, input interface{}) {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			parseValue(prefix+"_"+field.Name, fieldValue.Interface())
		default:
			name := fmt.Sprintf(prefix + "_" + field.Name)
			fmt.Println(name)
		}
	}
}

//func LoadEnvironment(input interface{}) {
//	//t := reflect.TypeOf(input)
//	//for i := 0; i < t.NumField(); i++ {
//	//	field := t.Field(i)
//	//	tag := field.Tag
//
//	//	fmt.Println("name:", field.Name)
//	//	fmt.Println("tag:", field.Tag)
//	//	fmt.Println("json:", strings.Split(tag.Get("json"), ","))
//	//}
//
//	t := reflect.ValueOf(input)
//	fmt.Println(t.Kind())
//
//	if t.Kind() == reflect.String {
//		if !t.CanSet() {
//			break
//		}
//
//		t.Set()
//	}
//
//	if t.Kind() == reflect.Struct {
//
//	}
//}
