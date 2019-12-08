package main

import (
	"fmt"
	"testing"
)

func TestConverter_HappyPath(t *testing.T) {
	input := "this is a sample string"
	expectedOutput := "this_is_a_sample_string"

	actual, err := converter(input)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if actual != expectedOutput {
		t.Errorf("expected output to be %s, but got %s", expectedOutput, actual)
	}
}

func TestConverter_ErrorPath(t *testing.T) {
	input := "12345678"
	expectedErr := fmt.Errorf("input length is eight which is forbidden")

	_, err := converter(input)
	if err == nil {
		t.Fatal("expected error not be nil, but was nil")
	}

	if err.Error() != expectedErr.Error() {
		t.Errorf("expected error to be %v, but got %v", expectedErr, err)
	}
}
