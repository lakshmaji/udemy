package cmd

import (
	reader_client "geektrust/clients/reader"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestCommands(t *testing.T) {

	tt := []struct {
		description string
		expected    [][]string
	}{
		{
			description: "single command",
			expected: [][]string{
				{"ADD_CERTIFICATION", "nADD_DEGREE"},
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
				"input.txt": {Data: []byte("ADD_CERTIFICATION 2\nADD_DEGREE 1")},
			}
			var reader reader_client.BaseReader = reader_client.New(fs)
			commandParser := New(reader)
			received, err := commandParser.Commands()

			if err != nil {
				t.Error("should not return error, received", err)
			}

			if reflect.DeepEqual(received, test.expected) {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
