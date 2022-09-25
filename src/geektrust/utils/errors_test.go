package utils

import "testing"

func TestErrorUnknownCategory(t *testing.T) {
	if ErrorUnknownCategory.Error() != "Unknown program category" {
		t.Error("Should be \"Unknown program category\"")
	}
}

func TestErrorNoFilePath(t *testing.T) {
	if ErrorNoFilePath.Error() != "Please provide the input file path" {
		t.Error("Should be \"Please provide the input file path\"")
	}
}

func TestErrorFileOpen(t *testing.T) {
	if ErrorFileOpen.Error() != "Error opening the input file" {
		t.Error("Should be \"Error opening the input file\"")
	}
}

func TestErrorUnknownCommand(t *testing.T) {
	if ErrorUnknownCommand != "Unrecognized command" {
		t.Error("Should be \"Unrecognized command\"")
	}
}

func TestUnknownCommandError(t *testing.T) {
	err := &UnknownCommandError{
		Command: "INCREASE_QUANTITY",
	}
	if err.Error() != "Unrecognized command INCREASE_QUANTITY" {
		t.Error("Should be \"Unrecognized command INCREASE_QUANTITY\"")
	}
}
