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
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
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
	var err error
	err = cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	err = cart.AddProgram(program.Program{
		Category: program.CategoryDiploma,
		Quantity: 1,
	})
	if err != nil {
		t.Error("Should not return error, received", err)
	}

	expectedNoOfPrograms := 3
	received := cart.TotalProgramsCount()
	if received != expectedNoOfPrograms {
		t.Errorf("Expected %d, Received %d", expectedNoOfPrograms, received)
	}
}

func TestAddProMembership(t *testing.T) {
	cart := &Cart{}
	cart.AddProMembership()
	if cart.ProMembershipFee() != ProMemberShipFee {
		t.Errorf("should have pro membership fee of %f", ProMemberShipFee)
	}
}

func TestTotalProMembershipDiscount(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	cart.AddProMembership()
	received := cart.TotalProMembershipDiscount()
	expectedDiscount := 300.0
	if received != expectedDiscount {
		t.Errorf("Expected %f, Received %f", expectedDiscount, received)
	}
}

func TestProgramsGrossAmount(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	received := cart.programsGrossAmount()
	expectedTotal := 10000.0
	if received != expectedTotal {
		t.Errorf("Expected %f, Received %f", expectedTotal, received)
	}
}

func TestProgramsNetAmount(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	received := cart.programsNetAmount()
	expectedTotal := 10000.0
	if received != expectedTotal {
		t.Errorf("Expected %f, Received %f", expectedTotal, received)
	}
}

func TestProgramsNetAmountWithProMembership(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	cart.AddProMembership()
	received := cart.programsNetAmount()
	expectedTotal := 9900.0
	if received != expectedTotal {
		t.Errorf("Expected %f, Received %f", expectedTotal, received)
	}
}

func TestZeroEnrollmentFee(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryDegree,
		Quantity: 2,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	received := cart.EnrollmentFee()
	expectedTotal := float64(0)
	if received != expectedTotal {
		t.Errorf("Expected %f, Received %f", expectedTotal, received)
	}
}

func TestEnrollmentFee(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryCertification,
		Quantity: 1,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	received := cart.EnrollmentFee()
	expectedTotal := float64(500)
	if received != expectedTotal {
		t.Errorf("Expected %f, Received %f", expectedTotal, received)
	}
}

func TestSubTotal(t *testing.T) {
	tt := []struct {
		description string
		cart        *Cart
		expected    float64
	}{
		{
			description: "program category degree",
			cart: &Cart{
				Programs: []program.Program{
					{Category: program.CategoryDegree, Quantity: 2},
				},
			},
			expected: 10000,
		},
		{
			description: "program category degree and pro-membership fee",
			cart: &Cart{
				Programs: []program.Program{
					{Category: program.CategoryDegree, Quantity: 2},
				},
				hasProMemberShip: true,
			},
			expected: 9900,
		},
		{
			description: "program category degree, pro-membership fee and enrollment fee",
			cart: &Cart{
				Programs: []program.Program{
					{Category: program.CategoryCertification, Quantity: 1},
				},
			},
			expected: 3500,
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := test.cart.SubTotal()
			if received != test.expected {
				t.Errorf("Expected %f, Received %f", test.expected, received)
			}
		})
	}

}

func TestTotal(t *testing.T) {
	cart := &Cart{}
	item := program.Program{
		Category: program.CategoryCertification,
		Quantity: 1,
	}
	err := cart.AddProgram(item)
	if err != nil {
		t.Error("Should not return error, received", err)
	}
	received := cart.Total()
	expectedTotal := float64(3500)
	if received != expectedTotal {
		t.Errorf("Expected %f, Received %f", expectedTotal, received)
	}
}
