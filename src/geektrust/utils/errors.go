package utils

import (
	"errors"
	"fmt"
)

// Error messages
var (
	ErrorUnknownCategory = errors.New("Unknown program category")
	ErrorNoFilePath      = errors.New("Please provide the input file path")
	ErrorFileOpen        = errors.New("Error opening the input file")
	ErrorUnknownCommand  = "Unrecognized command"
)

// UnknownCommandError - Each instantiation to UnknownCommandError returns a distinct error value even if the text is identical.
type UnknownCommandError struct {
	Command string
}

func (e *UnknownCommandError) Error() string {
	return fmt.Sprintf("%s: %s", ErrorUnknownCommand, e.Command)
}
