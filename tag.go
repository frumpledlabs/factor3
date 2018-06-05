package main

import (
	"errors"
	"regexp"
	"strings"
)

type tag struct {
	key    string
	values []string
}

func newTag(input string) (tag, error) {
	var t tag

	pair := strings.Split(input, ":")
	if len(pair) != 2 {
		return t, errors.New("Invalid tag: " + input)
	}

	key, err := parseKey(pair[0])
	if err != nil {
		return t, err
	}

	values, err := parseValues(pair[1])
	if err != nil {
		return t, err
	}

	t = tag{
		key:    key,
		values: values,
	}

	return t, nil
}

func parseKey(input string) (string, error) {
	key := strings.TrimSpace(input)

	if len(key) == 0 {
		return "", errors.New("Invalid key: " + input)
	}

	return key, nil
}

func parseValues(input string) ([]string, error) {
	values := []string{}

	input = strings.TrimSpace(input)
	matched, _ := regexp.MatchString(`"\S+"`, input)
	if !matched {
		return values, errors.New("Invalid tag values: " + input)
	}
	input = strings.Trim(input, "\"")

	for _, v := range strings.Split(input, ",") {
		if empty, _ := regexp.MatchString(v, `\s*`); empty {
			return values, errors.New("Empty value found")
		}
		values = append(values, v)
	}

	return values, nil
}

func (t tag) Key() string {
	return t.key
}

func (t tag) Values() []string {
	return t.values
}
