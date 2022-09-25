package core

import (
	"geektrust/core/program"
	"geektrust/utils"
	"testing"
)

func TestAddProgramNoError(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, Received", err)
	}
}

func TestAddProgram(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	cart.AddProgram(item)
	if cart.Programs[0].Category != program.CategoryDegree {
		t.Error("should contain DEGREE program category")
	}
}

func TestAddProgramError(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryUnknown,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err == nil {
		t.Error("Should return error", utils.ErrorUnknownCategory)
	}
}
