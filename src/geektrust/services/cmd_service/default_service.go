package cmd_service

import (
	reader_client "geektrust/clients/reader"
	"strings"
)

type service struct {
	reader reader_client.BaseReader
}

// New - creates Command parser given io.Reader implementation, which supports file system.
func New(reader reader_client.BaseReader) CommandParser {
	return &service{reader}
}

func (s *service) Commands() ([][]string, error) {
	name, err := s.reader.ParseFileName()
	if err != nil {
		return nil, err
	}

	// dirEntry, err := fs.ReadDir(os., file)
	content, err := s.reader.ParseFileContent(name)
	if err != nil {
		return nil, err
	}

	lines, err := s.reader.ParseFileLines(content)
	if err != nil {
		return nil, err
	}

	var commands [][]string

	for _, line := range lines {
		fields := strings.Fields(line)
		commands = append(commands, fields)
	}

	return commands, nil
}
