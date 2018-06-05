package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	output := toCliParameters(test{})
	fmt.Println(output)
}

type test struct {
	PublicFlag  string `flag:"--public-flag"`
	privateFlag string `flag:"--private-flag"`
	flagTest    string `a:"a",         b:"b",         c:"c"`
}

func toCliParameters(input interface{}) string {
	var buffer bytes.Buffer

	t := reflect.TypeOf(input)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println()

		field := t.Field(i)
		fmt.Printf("%+v\n", field)

		tag := field.Tag
		fmt.Printf("Tag:\t%+v\n\n", tag)
		fmt.Println(getTags(string(field.Tag)))
		fmt.Println()

		fmt.Println("name:", field.Name)
		fmt.Println("pkgPath:", field.PkgPath)
		fmt.Println("type:", field.Type)
		fmt.Println("tag:", field.Tag)
		fmt.Println("offset:", field.Offset)
		fmt.Println("index:", field.Index)
		fmt.Println("anonymous:", field.Anonymous)
	}

	output := strings.Join(strings.Fields(buffer.String()), " ")

	if output == "" || output == " " {
		return ""
	}

	return output
}

func getTags(tag string) []string {
	var tags []string
	var keys []string
	for _, t := range strings.Split(tag, ",") {
		tags = append(tags, strings.TrimSpace(t))
		keys = append(keys, strings.TrimSpace(strings.Split(t, ":")[0]))
	}

	fmt.Println("Keys:", keys)

	return tags
}
