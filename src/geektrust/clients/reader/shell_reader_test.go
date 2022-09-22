package reader

import (
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

func TestParseFileContent(t *testing.T) {

	// Mock os.Open
	originalOsOpen := osOpen
	defer func() { osOpen = originalOsOpen }()

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
		"input.txt":       {Data: []byte("ADD_CERTIFICATION 2\nADD_DEGREE 1")},
	}

	// mockOpen := func(name string) (*os.File, error) {
	// 	// file := os.Stdin
	// 	// file.WriteString("sd\n")
	// 	// file.WriteString("sd\n")
	// 	// file, err := os.Create("./temp  .txt")
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	// writer := bufio.NewWriterSize(file, 10)
	// 	// linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
	// 	// for _, line := range linesToWrite {
	// 	// 	_, err := writer.WriteString(line + "\n")
	// 	// 	if err != nil {
	// 	// 		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	// 	// 	}
	// 	// }
	// 	fstest.TestFS{}
	// 	var file bytes.Buffer
	// 	file.WriteString("dasdsa")
	// 	return file, nil
	// }
	// osOpen = mockOpen

	mockReader := New(fs)
	// e := fstest.TestFS()
	// dirEntry, err := fs.ReadDir("input.txt")

	content, err := mockReader.ParseFileContent("input.txt")
	if err != nil {
		t.Errorf("should not throw error, got %v", err)
	}
	// lines, err := mockReader.ReadLines(fileInput)
	// if err != nil {
	// 	t.Errorf("should not throw error, got %v", err)
	// }
	// t.Log(lines)
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

	// defer content.(*os.File).Close()
	// scanner := bufio.NewScanner(content)
	// var lines []string

	// for scanner.Scan() {
	// 	textLine := scanner.Text()
	// 	lines = append(lines, textLine)
	// }

	// if len(lines) != 0 {
	// 	t.Error("hsgdhs")

	// }
	// rd := bufio.NewReader(content)
	// str, _ := rd.ReadString("\n")
	// io.WriteString(out, strings.TrimSuffix(str, "\n")+" was input\n")

	// if receivedFileName != "input.txt" {
	// 	t.Errorf("Expected file name %s, received %s", "input.txt", receivedFileName)
	// }
}

// func TestFileInputError(t *testing.T) {
// 	str := ("Package Id, Discount, Total Delivery Cost\n")
// 	str += ("pkg1, 0.00, 175.00\n")
// 	str += ("pkg2, 0.00, 275.00\n")
// 	str += ("pkg3, 35.00, 665.00")

// 	expected := "Package Id, Discount, Total Delivery Cost\npkg1, 0.00, 175.00\npkg2, 0.00, 275.00\npkg3, 35.00, 665.00\n"
// 	var output bytes.Buffer
// 	New(&output, DefaultOptions).WriteLn(str)

// 	if output.String() != expected {
// 		t.Errorf("Expected %v, got %v", str, output)
// 	}
// }
