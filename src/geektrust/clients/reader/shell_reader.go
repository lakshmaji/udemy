package reader

import (
	"bufio"
	"geektrust/utils"
	"io"
	"os"
)

var osOpen = os.Open
var osArgs = os.Args

// Shell reader client
type client struct {
}

// This reader is shell based, which means it can interact with local file system.
//
// Currently it indirectly support this application by providing interface to read data from file.
func New() BaseReader {
	return &client{}
}

func (f *client) ParseFileName() (string, error) {
	args := osArgs[1:]
	if len(args) == 0 {
		return "", utils.ErrorNoFilePath
	}
	return args[0], nil
}

func (f *client) ParseFileContent(fileSystem FileSystem, name string) (io.Reader, error) {
	file, err := fileSystem.Open(name)
	if err != nil {
		return nil, utils.ErrorFileOpen
	}
	return file, nil
}

func (f *client) ParseFileLines(file io.Reader) ([]string, error) {
	defer file.(*os.File).Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textLine := scanner.Text()
		lines = append(lines, textLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
