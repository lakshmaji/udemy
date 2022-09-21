package clients

import (
	"bufio"
	"errors"
	"os"
)

type fileClient struct {
}

// Handles responsibility of reading from **STDIN**
func NewCartFileReader() BaseReader {
	return &fileClient{}
}

func (f *fileClient) FileInput() ([]string, error) {
	cliArgs := os.Args[1:]
	if len(cliArgs) == 0 {
		return nil, errors.New("Please provide the input file path")
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("Error opening the input file")
	}

	defer file.Close()

	// var commands [][]string
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	textLine := scanner.Text()
	// 	fields := strings.Fields(textLine)
	// 	commands = append(commands, fields)
	// }
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
