package main

import (
	"fmt"
	"testing"
)

func TestConverter_ErrorPath(t *testing.T) {
	type test struct {
		name           string
		input          string
		expectedOutput string
		expectedErr    error
	}

	tests := []test{
		test{
			name:           "simple input",
			input:          "this is a sample string",
			expectedOutput: "this_is_a_sample_string",
			expectedErr:    nil,
		},
		test{
			name:           "complex input with special characters",
			input:          "this@#$@#$is a complex!!!!string...",
			expectedOutput: "this_is_a_complex_string",
			expectedErr:    nil,
		},
		test{
			name:        "unhappy path, invalid length",
			input:       "12345678",
			expectedErr: fmt.Errorf("input length is eight which is forbidden"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := converter(test.input)

			if test.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected an error, but got nil")
				}

				if test.expectedErr.Error() != err.Error() {
					t.Fatalf("expected err to be %s, but got %s", test.expectedErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, but got %s", err)
			}

			if test.expectedOutput != actual {
				t.Errorf("expected output to be %s, but got %s", test.expectedOutput, actual)
			}
		})
	}
}
