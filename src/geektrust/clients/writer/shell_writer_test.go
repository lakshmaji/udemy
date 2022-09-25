package writer

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {

	tt := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			description: "TestWrite",
			input:       "Hello World",
			expected:    "Hello World\n",
		},
		{
			description: "TestWrite",
			input:       fmt.Sprintf("%d", 10),
			expected:    "10\n",
		},
		{
			description: "TestWrite",
			input:       fmt.Sprintf("%t", true),
			expected:    "true\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var output bytes.Buffer
			New(&output, DefaultOptions).WriteLn(tc.input)
			if output.String() != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, output)
			}
		})
	}
}

func TestWriteWithMultiLineString(t *testing.T) {

	tt := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			description: "TestWriteWithMultiLineString",
			input:       "Hello World\nHello World",
			expected:    "Hello World\nHello World\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var output bytes.Buffer
			New(&output, DefaultOptions).WriteLn(tc.input)
			if output.String() != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, output)
			}
		})
	}
}

func TestWriteError(t *testing.T) {

	tt := []struct {
		description string
		input       interface{}
		expected    interface{}
	}{
		{
			description: "TestWrite",
			input:       "Hello Error",
			expected:    "Hello Error",
		},
		{
			description: "TestWrite",
			input:       10,
			expected:    "10",
		},
		{
			description: "TestWrite",
			input:       true,
			expected:    "true",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var output bytes.Buffer
			defer func() {
				r := recover()
				if output.String() != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, output)
				}
				if r == nil {
					t.Errorf("Should panic")
				}
			}()
			New(&output, DefaultTestOptions).WriteError(tc.input)
			if output.String() != "" {
				t.Errorf("Should throw error, received %v", output)
			}
		})
	}
}

func TestExit(t *testing.T) {

	var spy string
	// mock
	exitProgram = func(exitCode int) {
		spy = fmt.Sprintf("Exiting with code %d", exitCode)
	}
	input := "Some error string"

	var output bytes.Buffer
	defer func() {
		r := recover()
		t.Log(spy)

		if spy != "Exiting with code 1" {
			t.Error("Should exit with code 1")
		}
		if r != nil {
			t.Error("Should not return any error")
		}
		// restore
		exitProgram = os.Exit
	}()
	New(&output, DefaultOptions).WriteError(input)
}
