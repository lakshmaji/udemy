package cart_service

import (
	"geektrust/core"
	cart_model "geektrust/core"
	"geektrust/core/coupon"
	"geektrust/core/program"
	"geektrust/utils"
	"reflect"
	"strconv"
	"testing"
)

func TestAddCoupon(t *testing.T) {

	tt := []struct {
		description string
		input       []string
		expected    coupon.Coupons
	}{
		{
			description: "invalid coupon added",
			input:       []string{"DEAL_NEW_ACCOUNT"},
			expected:    coupon.Coupons{coupon.Coupon("DEAL_NEW_ACCOUNT")},
		},
		{
			description: "valid coupon added",
			input:       []string{"DEAL_G20"},
			expected:    coupon.Coupons{coupon.Coupon("DEAL_G20")},
		},
		{
			description: "add multiple coupons",
			input:       []string{"DEAL_G20", "B4G1"},
			expected:    coupon.Coupons{coupon.Coupon("DEAL_G20"), coupon.Coupon("B4G1")},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			cart := &core.Cart{}
			mockCouponService := mockNewCouponService(mockInput{applicableCoupon: "", discount: 0})
			cartService := New(cart, mockCouponService)
			for _, code := range test.input {
				cartService.AddCoupon(code)
			}
			received := cart.CouponsList
			if !reflect.DeepEqual(received, test.expected) {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}

}

func TestAddProMembership(t *testing.T) {

	tt := []struct {
		description string
		invoke      bool
		expected    float64
	}{
		{
			description: "not subscribed to pro-membership",
			invoke:      false,
			expected:    0,
		},
		{
			description: "subscribed to pro-membership",
			invoke:      true,
			expected:    cart_model.ProMemberShipFee,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			cart := &core.Cart{}
			mockCouponService := mockNewCouponService(mockInput{applicableCoupon: "", discount: 0})
			cartService := New(cart, mockCouponService)
			if test.invoke {
				cartService.AddProMembership()
			}
			received := cart.ProMembershipFee()
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}

}

func TestAddProgram(t *testing.T) {

	expected := program.CategoryDegree

	cart := &core.Cart{}
	mockCouponService := mockNewCouponService(mockInput{applicableCoupon: "", discount: 0})
	cartService := New(cart, mockCouponService)

	err := cartService.AddProgram("2", "DEGREE")
	if err != nil {
		t.Errorf("Should not return error, Received %v", err)
	}

	if int(cart.Programs[0].Category) != int(expected) {
		t.Errorf("Expected %v, Received %v", expected, cart.Programs[0].Category)
	}

}

func TestAddProgramError(t *testing.T) {
	const fnAtoi = "Atoi"

	type input struct {
		quantity string
		category string
	}

	tt := []struct {
		description string
		input       input
		expected    error
	}{
		{
			description: "invalid quantity",
			input: input{
				quantity: "one",
				category: "ADD_DEGREE",
			},
			expected: &strconv.NumError{Func: fnAtoi, Num: "one", Err: strconv.ErrSyntax},
		},
		{
			description: "invalid category",
			input: input{
				quantity: "1",
				category: "CREATE_DEGREE",
			},
			expected: utils.ErrorUnknownCategory,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			cart := &core.Cart{}
			mockCouponService := mockNewCouponService(mockInput{applicableCoupon: "", discount: 0})
			cartService := New(cart, mockCouponService)

			err := cartService.AddProgram(test.input.quantity, test.input.category)
			if err == nil {
				t.Error("Should return error")
			}
			if err.Error() != test.expected.Error() {
				t.Errorf("Expected %v, Received %v", test.expected, err)
			}
		})
	}

}

func TestComputeDiscount(t *testing.T) {

	tt := []struct {
		description        string
		input              cart_model.Cart
		expectedCoupon     coupon.Coupon
		expectedDiscount   float64
		couponSvcMockInput mockInput
	}{
		{
			description: "no coupons",
			input: cart_model.Cart{
				Programs: []program.Program{
					{
						Category: program.CategoryDegree,
						Quantity: 2,
					},
				},
			},
			expectedCoupon:     coupon.Coupon(""),
			expectedDiscount:   0,
			couponSvcMockInput: mockInput{applicableCoupon: "", discount: 0},
		},
		{
			description: "coupon applied",
			input: cart_model.Cart{
				Programs: []program.Program{
					{
						Category: program.CategoryDegree,
						Quantity: 4,
					},
				},
				CouponsList: coupon.Coupons{coupon.CouponDealG20},
			},
			expectedCoupon:     coupon.CouponDealG20,
			expectedDiscount:   0,
			couponSvcMockInput: mockInput{applicableCoupon: "DEAL_G20", discount: 0},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			mockCouponService := mockNewCouponService(test.couponSvcMockInput)
			cart := &test.input
			cartService := New(&test.input, mockCouponService)

			cartService.ComputeDiscount()
			if cart.CouponApplied != test.expectedCoupon {
				t.Errorf("Expected %v, Received %v", test.expectedCoupon, cart.CouponApplied)
			}
			if cart.CouponDiscountApplied != test.expectedDiscount {
				t.Errorf("Expected %v, Received %v", test.expectedDiscount, cart.CouponDiscountApplied)
			}
		})
	}

}
