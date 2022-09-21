package command_parser

import (
	reader_client "geektrust/clients/reader"
	"strings"
)

type service struct {
	reader reader_client.BaseReader
}

// Command parser for shell client input
func New(reader reader_client.BaseReader) CommandParser {
	return &service{reader}
}

func (s *service) Commands() ([][]string, error) {
	lines, err := s.reader.FileInput()
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
