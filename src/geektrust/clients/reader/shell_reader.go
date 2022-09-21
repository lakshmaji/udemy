package reader

import (
	"bufio"
	"geektrust/utils"
	"os"
)

// Shell reader client
type client struct {
}

// This reader is shell based, which means it can interact with local file system.
//
// Currently it indirectly support this application by providing interface to read data from file.
func New() BaseReader {
	return &client{}
}

func (f *client) FileInput() ([]string, error) {
	cliArgs := os.Args[1:]
	if len(cliArgs) == 0 {
		return nil, utils.ErrorNoFilePath
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		return nil, utils.ErrorFileOpen
	}

	defer file.Close()

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
