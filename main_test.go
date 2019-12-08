package main

import "testing"

func TestConverter(t *testing.T) {
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
