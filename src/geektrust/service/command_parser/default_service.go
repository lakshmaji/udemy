package command_parser

import (
	"geektrust/clients"
	"strings"
)

type service struct {
	reader clients.BaseReader
}

// Command parser for shell client input
func NewShellCommandParser(reader clients.BaseReader) CommandParser {
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
