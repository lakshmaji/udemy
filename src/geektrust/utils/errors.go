package utils

import (
	"errors"
	"fmt"
)

var (
	ErrorUnknownCategory = errors.New("Unknown program category")
	ErrorNoFilePath      = errors.New("Please provide the input file path")
	ErrorFileOpen        = errors.New("Error opening the input file")
	ErrorUnknownCommand  = "Unrecognized command"
)

type UnknownCommandError struct {
	Command string
}

func (e *UnknownCommandError) Error() string {
	return fmt.Sprintf("%s: %s", ErrorUnknownCommand, e.Command)
}
