package reader

import (
	"geektrust/utils"
	"io"
	"testing"
	"testing/fstest"
)

// func TestFileInput(t *testing.T) {

// 	// Mock os.Open
// 	originalOsOpen := osOpen
// 	defer func() { osOpen = originalOsOpen }()
// 	var receivedFileName string
// 	mockOpen := func(name string) (*os.File, error) {
// 		receivedFileName = name
// 		// return nil, &os.PathError{Op: "open", Path: name, Err: errors.New("Some error")}
// 		file := os.File{}
// 		// file := os.NewFile()
// 		// buf := new(bytes.Buffer)
// 		// file := buf.Bytes()
// 		// file.Seek(int64(os.SEEK_CUR), os.SEEK_CUR)
// 		// file.WriteString("ADD_DEGREE 2")
// 		return &file, nil
// 	}
// 	osOpen = mockOpen

// 	// Mock os.Args
// 	originalOsArgs := osArgs
// 	defer func() { osArgs = originalOsArgs }()

// 	mockArgs := []string{"main.go", "input.txt"}
// 	osArgs = mockArgs

// 	lines, err := New().FileInput()
// 	if err != nil {
// 		t.Errorf("should not throw error, got %v", err)
// 	}
// 	t.Log(lines)
// 	if receivedFileName == "" {
// 		t.Error("File name should not be empty")
// 	}

// 	if receivedFileName != "input.txt" {
// 		t.Errorf("Expected file name %s, received %s", "input.txt", receivedFileName)
// 	}
// }

// func TestParseFileContent(t *testing.T) {

// 	// Mock os.Open
// 	// originalOsOpen := osOpen
// 	// defer func() { osOpen = originalOsOpen }()

// 	fs := fstest.MapFS{
// 		"input.txt": {Data: []byte("ADD_CERTIFICATION 2\nADD_DEGREE 1")},
// 	}

// 	// mockOpen := func(name string) (*os.File, error) {
// 	// 	// file := os.Stdin
// 	// 	// file.WriteString("sd\n")
// 	// 	// file.WriteString("sd\n")
// 	// 	// file, err := os.Create("./temp  .txt")
// 	// 	// if err != nil {
// 	// 	// 	log.Fatal(err)
// 	// 	// }
// 	// 	// writer := bufio.NewWriterSize(file, 10)
// 	// 	// linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
// 	// 	// for _, line := range linesToWrite {
// 	// 	// 	_, err := writer.WriteString(line + "\n")
// 	// 	// 	if err != nil {
// 	// 	// 		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
// 	// 	// 	}
// 	// 	// }
// 	// 	fstest.TestFS{}
// 	// 	var file bytes.Buffer
// 	// 	file.WriteString("dasdsa")
// 	// 	return file, nil
// 	// }
// 	// osOpen = mockOpen

// 	mockReader := New(fs)
// 	// e := fstest.TestFS()
// 	// dirEntry, err := fs.ReadDir("input.txt")

// 	content, err := mockReader.ParseFileContent("input.txt")
// 	if err != nil {
// 		t.Errorf("should not throw error, got %v", err)
// 	}
// 	// lines, err := mockReader.ReadLines(fileInput)
// 	// if err != nil {
// 	// 	t.Errorf("should not throw error, got %v", err)
// 	// }
// 	// t.Log(lines)
// 	if content == nil {
// 		t.Error("File content should not be empty")
// 	}

// 	b, err := io.ReadAll(content)
// 	if err != nil {
// 		t.Error("error reading file", err)
// 	}
// 	value := string(b)
// 	if len(value) == 0 {
// 		t.Error("No Content available in the file", value)

// 	}

// 	// defer content.(*os.File).Close()
// 	// scanner := bufio.NewScanner(content)
// 	// var lines []string

// 	// for scanner.Scan() {
// 	// 	textLine := scanner.Text()
// 	// 	lines = append(lines, textLine)
// 	// }

// 	// if len(lines) != 0 {
// 	// 	t.Error("hsgdhs")

// 	// }
// 	// rd := bufio.NewReader(content)
// 	// str, _ := rd.ReadString("\n")
// 	// io.WriteString(out, strings.TrimSuffix(str, "\n")+" was input\n")

// 	// if receivedFileName != "input.txt" {
// 	// 	t.Errorf("Expected file name %s, received %s", "input.txt", receivedFileName)
// 	// }
// }

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

// func TestParseFileLines(t *testing.T) {

// 	fs := fstest.MapFS{}
// 	mockReader := New(fs)

// 	// var file bytes.Buffer
// 	// file.WriteString("dasdsa")
// 	a, _ := fs.Open("input.txt")

// 	lines, err := mockReader.ParseFileLines(a)
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
