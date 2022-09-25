package cmd_service

import (
	"errors"
	"geektrust/utils"
	"reflect"
	"strings"
	"testing"
)

func TestCommands(t *testing.T) {

	tt := []struct {
		description string
		expected    [][]string
		input       mockInput
	}{
		{
			description: "single line command",
			input: mockInput{
				FileName:    "input.txt",
				FileContent: strings.NewReader("ADD_CERTIFICATION\t2"),
				FileLines:   []string{"ADD_CERTIFICATION\t2"},
			},
			expected: [][]string{
				{"ADD_CERTIFICATION", "2"},
			},
		},
		{
			description: "multi line command",
			input: mockInput{
				FileName:    "input.txt",
				FileContent: strings.NewReader("ADD_CERTIFICATION 2\nADD_DEGREE 1\nPRINT_BILL"),
				FileLines:   []string{"ADD_CERTIFICATION 2\n", "ADD_DEGREE 1", "PRINT_BILL"},
			},
			expected: [][]string{
				{"ADD_CERTIFICATION", "2"},
				{"ADD_DEGREE", "1"},
				{"PRINT_BILL"},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			commandParser := New(mockNewReader(test.input))
			received, err := commandParser.Commands()

			if err != nil {
				t.Error("should not return error, received", err)
			}

			if !reflect.DeepEqual(received, test.expected) {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}

func TestCommandsError(t *testing.T) {

	ErrUnableToParseLines := errors.New("Unable to parse commands")

	type fileContent struct {
		Data []byte
	}

	tt := []struct {
		description string
		expected    error
		input       mockInput
	}{
		{
			description: "no file name was provided",
			input: mockInput{
				ErrParseFileName: utils.ErrorNoFilePath,
			},
			expected: utils.ErrorNoFilePath,
		},
		{
			description: "no file content was provided",
			input: mockInput{
				FileName:            "input.txt",
				ErrParseFileContent: utils.ErrorFileOpen,
			},
			expected: utils.ErrorFileOpen,
		},
		{
			description: "error while reading line by line",
			input: mockInput{
				FileName:          "input.txt",
				FileContent:       strings.NewReader("ADD_CERTIFICATION"),
				ErrParseFileLines: ErrUnableToParseLines,
			},
			expected: ErrUnableToParseLines,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			commandParser := New(mockNewReader(test.input))
			received, err := commandParser.Commands()

			if err == nil {
				t.Error("should return error, received", err)
			}

			if test.expected != err {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}

			if len(received) > 0 {
				t.Errorf("Should not return commands, Received %v", received)
			}
		})
	}
}
