package main

import (
	"errors"
	"regexp"
	"strings"
)

type tag struct {
	key   string
	value string
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

	value, err := parseValue(pair[1])
	if err != nil {
		return t, err
	}

	t = tag{
		key:   key,
		value: value,
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

func parseValue(input string) (string, error) {
	input = strings.TrimSpace(input)
	matched, _ := regexp.MatchString("\".*\"", input)
	if !matched {
		return "", errors.New("Invalid tag value: " + input)
	}
	input = strings.Trim(input, "\"")

	return input, nil
}

func (t tag) Key() string {
	return t.key
}

func (t tag) Value() string {
	return t.value
}
