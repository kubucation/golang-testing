package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(converter("this is some####sample string!@#4 you"))
}

func converter(input string) (string, error) {
	if len(input) == 7 {
		return "", fmt.Errorf("input length is seven which is forbidden")
	}

	splitter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	splits := strings.FieldsFunc(input, splitter)
	out := strings.Join(splits, "_")

	return out, nil
}
