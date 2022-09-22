package utils

import "errors"

var (
	ErrorUnknownCategory = errors.New("Unknown program category")
	ErrorNoFilePath      = errors.New("Please provide the input file path")
	ErrorFileOpen        = errors.New("Error opening the input file")
	ErrorUnknownCommand  = "Unrecognized command"
)
