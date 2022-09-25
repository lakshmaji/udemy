package cmd

import (
	reader_client "geektrust/clients/reader"
	"geektrust/utils"
	"reflect"
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

func TestCommandsError(t *testing.T) {

	type fileContent struct {
		Data []byte
	}

	tt := []struct {
		description string
		files       []string
		fileContent fstest.MapFS
		expected    error
	}{
		{
			description: "no file name was provided",
			files:       []string{"main.go"},
			fileContent: fstest.MapFS{},
			expected:    utils.ErrorNoFilePath,
		},
		{
			description: "no file content was provided",
			files:       []string{"main.go", "input.txt"},
			fileContent: fstest.MapFS{
				"ouput.txt": {Data: []byte("")},
			},
			expected: utils.ErrorFileOpen,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			// mock os.Args
			originalOsArgs := reader_client.OsArgs
			defer func() { reader_client.OsArgs = originalOsArgs }()

			mockArgs := test.files
			reader_client.OsArgs = mockArgs

			// mock fs
			fs := test.fileContent
			var reader reader_client.BaseReader = reader_client.New(fs)
			commandParser := New(reader)
			received, err := commandParser.Commands()

			if err == nil {
				t.Error("should return error, received", err)
			}

			if test.expected != err {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}

			if len(received) != 0 {
				t.Errorf("Should not return commands, Received %v", received)
			}
		})
	}
}
