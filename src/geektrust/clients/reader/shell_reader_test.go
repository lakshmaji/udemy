package reader

import (
	"geektrust/utils"
	"io"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

func TestParseFileContent(t *testing.T) {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte("ADD_CERTIFICATION 2\nADD_DEGREE 1")},
	}

	mockReader := New(fs)
	content, err := mockReader.ParseFileContent("input.txt")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
	if content == nil {
		t.Error("File content should not be empty")
	}

	b, err := io.ReadAll(content)
	if err != nil {
		t.Error("error reading file", err)
	}
	value := string(b)
	if len(value) == 0 {
		t.Error("No Content available in the file", value)
	}
}

func TestParseFileContentError(t *testing.T) {
	fs := fstest.MapFS{}

	mockReader := New(fs)
	content, err := mockReader.ParseFileContent("input.txt")
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}
	if content != nil {
		t.Error("File content is empty")
	}
	if err != utils.ErrorFileOpen {
		t.Errorf("Expected %v, Received %v", utils.ErrorFileOpen, err)
	}
}

func TestParseFileNameError(t *testing.T) {
	// mock
	originalOsArgs := OsArgs
	defer func() { OsArgs = originalOsArgs }()

	// restore
	mockArgs := []string{"main.go"}
	OsArgs = mockArgs

	fs := fstest.MapFS{}
	reader := New(fs)

	_, err := reader.ParseFileName()
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}
}

func TestParseFileName(t *testing.T) {
	// mock
	originalOsArgs := OsArgs
	defer func() { OsArgs = originalOsArgs }()

	// restore
	mockArgs := []string{"main.go", "input.txt"}
	OsArgs = mockArgs

	fs := fstest.MapFS{}
	reader := New(fs)
	filename, err := reader.ParseFileName()
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
	if filename != "input.txt" {
		t.Error("Should return input.txt as filename")
	}
}

// func TestParseFileLinesError(t *testing.T) {

// 	fs := fstest.MapFS{}
// 	reader := New(fs)
// 	a, _ := fs.Open("input.txt")

// 	lines, err := reader.ParseFileLines(a)
// 	if err != nil {
// 		t.Errorf("should not return error, got %v", err)
// 	}
// 	expectedLines := [][]string{
// 		{
// 			"ADD_CERTIFICATION", "2",
// 		},
// 	}
// 	if !reflect.DeepEqual(lines, expectedLines) {
// 		t.Errorf("Expected %v, Received %v", expectedLines, lines)
// 	}
// }

func TestParseFileLines(t *testing.T) {

	var builder strings.Builder
	builder.WriteString("ADD_CERTIFICATION\t2\n")
	builder.WriteString("ADD_DEGREE\t1\n")

	mfs := fstest.MapFS{
		"input.txt": {Data: []byte(builder.String())},
	}
	mockReader := New(mfs)
	content, _ := mfs.Open("input.txt")

	lines, err := mockReader.ParseFileLines(content)
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
	expectedLines := []string{
		"ADD_CERTIFICATION\t2",
		"ADD_DEGREE\t1",
	}
	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("Expected %v, Received %v", expectedLines, lines)
	}
}
