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
	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte("ADD_CERTIFICATION 2\nADD_DEGREE 1")},
	}

	shellReader := New(mockFS)
	file, err := shellReader.ParseFileContent("input.txt")
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
	if file == nil {
		t.Error("File content should not be empty")
	}

	content, err := io.ReadAll(file)
	if err != nil {
		t.Error("error reading file", err)
	}
	value := string(content)
	if len(value) == 0 {
		t.Error("No Content available in the file", value)
	}
}

func TestParseFileContentError(t *testing.T) {
	mockFS := fstest.MapFS{}

	shellReader := New(mockFS)
	content, err := shellReader.ParseFileContent("input.txt")
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
	mockArgs := []string{"main.go"}
	OsArgs = mockArgs

	mockFS := fstest.MapFS{}
	reader := New(mockFS)

	_, err := reader.ParseFileName()
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}
}

func TestParseFileName(t *testing.T) {
	// mock
	originalOsArgs := OsArgs
	defer func() { OsArgs = originalOsArgs }()
	mockArgs := []string{"main.go", "input.txt"}
	OsArgs = mockArgs

	mockFS := fstest.MapFS{}
	reader := New(mockFS)
	filename, err := reader.ParseFileName()
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
	if filename != "input.txt" {
		t.Error("Should return input.txt as filename")
	}
}

func TestParseFileNameAbs(t *testing.T) {
	// mock
	originalOsArgs := OsArgs
	defer func() { OsArgs = originalOsArgs }()
	mockArgs := []string{"main.go", "/input.txt"}
	OsArgs = mockArgs

	mockFS := fstest.MapFS{}
	reader := New(mockFS)
	filename, err := reader.ParseFileName()
	if err != nil {
		t.Errorf("should not return error, got %v", err)
	}
	if filename != "input.txt" {
		t.Error("Should return input.txt as filename")
	}
}

func TestParseFileLines(t *testing.T) {

	var builder strings.Builder
	builder.WriteString("ADD_CERTIFICATION\t2\n")
	builder.WriteString("ADD_DEGREE\t1\n")

	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte(builder.String())},
	}
	mockReader := New(mockFS)
	content, _ := mockFS.Open("input.txt")

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
