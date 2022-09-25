package cmd

import (
	"errors"
	reader_client "geektrust/clients/reader"
	"geektrust/utils"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

func TestCommands(t *testing.T) {

	tt := []struct {
		description string
		input       []byte
		expected    [][]string
	}{
		{
			description: "single line command",
			input:       []byte("ADD_CERTIFICATION 2"),
			expected: [][]string{
				{"ADD_CERTIFICATION", "2"},
			},
		},
		{
			description: "multi line command",
			input:       []byte("ADD_CERTIFICATION 2\nADD_DEGREE 1\nPRINT_BILL"),
			expected: [][]string{
				{"ADD_CERTIFICATION", "2"},
				{"ADD_DEGREE", "1"},
				{"PRINT_BILL"},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			// mock os.Args
			originalOsArgs := reader_client.OsArgs
			defer func() { reader_client.OsArgs = originalOsArgs }()

			mockArgs := []string{"main.go", "input.txt"}
			reader_client.OsArgs = mockArgs

			// mock fs
			fs := fstest.MapFS{
				"input.txt": {Data: test.input},
			}
			var reader reader_client.BaseReader = reader_client.New(fs)
			commandParser := New(reader)
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

var (
	ErrUnableToParseLines = errors.New("Unable to parse commands")
)

func TestCommandsError(t *testing.T) {

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
			commandParser := New(newMock(test.input))
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
