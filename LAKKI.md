# Instructions

## Development

```bash
go run main.go sample_input/input1.txt
```

## Testing

```bash
$HOME/go/bin/richgo test -coverprofile=c.out ./... 
go tool cover -html=c.out 
```

## Lint


```bash
go fmt ./...
go vet
go vet ./...
$HOME/go/bin/golint ./...

 $HOME/go/bin/golangci-lint run
 $HOME/go/bin/gocyclo .   
```

![golint,gofmt, go vet](https://sparkbox.com/uploads/article_uploads/code-checking-options.png)


## Code snippets

### File open mock service


```go
package reader

import (
	"io/fs"
	"os"
	"testing/fstest"
)

type fileSystemMock struct {
}

// FIXME: not using this any where. This can be removed.
func MockNewFile() fs.FS {
	return &fileSystemMock{}
}

func (f fileSystemMock) Open(name string) (fs.File, error) {
	const (
		firstBody  = "Post 1\nDescription"
		secondBody = "Post 2 \n Description"
	)

	mfs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
		"input.txt":       {Data: []byte("ADD BANANA")},
	}

	dir, err := fs.ReadDir(mfs, ".")
	if err != nil {
		return nil, err
	}

	for _, f := range dir {
		// post, err := getPost(fileSystem, f)
		postFile, err := mfs.Open(f.Name())

		if err != nil {
			//todo: needs clarification, should we totally fail if one file fails? or just ignore?
			return nil, err
		}
		defer postFile.Close()
		return postFile, nil
	}

	return nil, &fs.PathError{
		Op:   "read",
		Path: name,
		Err:  os.ErrNotExist,
	}
}

```



```go
package cmd_service

import (
	"errors"
	"geektrust/utils"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
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
				FileName: "input.txt",
				// FileContent: strings.NewReader("ADD_CERTIFICATION\t2"),
				FileContent: func() fs.File {
					fs := fstest.MapFS{
						"input.txt": {Data: []byte("ADD_CERTIFICATION 2")},
					}
					a, _ := fs.Open("input.txt")
					return a
				}(),
				FileLines: []string{"ADD_CERTIFICATION\t2"},
			},
			expected: [][]string{
				{"ADD_CERTIFICATION", "2"},
			},
		},
		{
			description: "multi line command",
			input: mockInput{
				FileName: "input.txt",
				// FileContent: strings.NewReader("ADD_CERTIFICATION 2\nADD_DEGREE 1\nPRINT_BILL"),
				FileContent: func() fs.File {
					fs := fstest.MapFS{
						"input.txt": {Data: []byte("ADD_CERTIFICATION 2\nADD_DEGREE\t1\nPRINT_BILL")},
					}
					a, _ := fs.Open("input.txt")
					return a
				}(),
				FileLines: []string{"ADD_CERTIFICATION 2\n", "ADD_DEGREE 1", "PRINT_BILL"},
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
				FileName: "input.txt",
				// FileContent:       strings.NewReader("ADD_CERTIFICATION"),
				FileContent: func() fs.File {
					fs := fstest.MapFS{
						"input.txt": {Data: []byte("ADD_CERTIFICATION")},
					}
					a, _ := fs.Open("input.txt")
					return a
				}(),
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

```