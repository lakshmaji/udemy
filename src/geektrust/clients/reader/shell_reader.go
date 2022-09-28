package reader

import (
	"bufio"
	"geektrust/utils"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// OsArgs hold the reference to command-line arguments.
// This allows for mocking during testing.
var OsArgs = os.Args

// Shell reader client
type client struct {
	fileSystem FileSystem
}

// New reader allows reading from shell, which means it can interact with local file system.
//
// Currently it indirectly support this application by providing interface to read data from file.
func New(fileSystem FileSystem) BaseReader {
	return &client{fileSystem}
}

func (f *client) ParseFileName() (string, error) {
	args := OsArgs[1:]
	if len(args) == 0 {
		return "", utils.ErrorNoFilePath
	}
	// input_path := args[0]
	// if filepath.IsAbs(input_path) {
	// 	// unroot, as root is defined by DirFS
	// 	return input_path[1:], nil
	// }
	// return input_path, nil
	return filepath.Base(args[0]), nil
}

func (f *client) ParseFileContent(name string) (io.Reader, error) {
	file, err := f.fileSystem.Open(name)
	if err != nil {
		return nil, utils.ErrorFileOpen
	}
	return file, nil
}

func (f *client) ParseFileLines(file io.Reader) ([]string, error) {
	defer file.(fs.File).Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textLine := scanner.Text()
		lines = append(lines, textLine)
	}

	return lines, nil
}
