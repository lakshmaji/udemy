package core

import (
	"geektrust/core/coupon"
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

func TestAddCoupon(t *testing.T) {
	cart := &Cart{}
	cart.AddCoupon(coupon.CouponB4G1)
	if cart.CouponsList[0] != coupon.CouponB4G1 {
		t.Error("should contain B4G1 coupon code")
	}
}

func TestTotalProgramsCount(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	cart.AddProgram(item)
	cart.AddProgram(program.Program{
		Category: program.CategoryDiploma,
		Quantity: 1,
	})

	expected := 3
	received := cart.TotalProgramsCount()
	if received != expected {
		t.Errorf("Expected %d, Received %d", expected, received)
	}
}
