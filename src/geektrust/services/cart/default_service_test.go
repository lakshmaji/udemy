package cart

import (
	"geektrust/core"
	cart_model "geektrust/core"
	"geektrust/core/coupon"
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

// func TestAddProgram(t *testing.T) {

// 	type input struct {
// 		quantity string
// 		category string
// 	}
// 	tt := []struct {
// 		description string
// 		input       input
// 		expected    float64
// 	}{
// 		{
// 			description: "invalid quantity",
// 			input: input{
// 				quantity: "one",
// 				category: "ADD_DEGREE",
// 			},
// 			expected: 0,
// 		},
// 	}

// 	for _, test := range tt {
// 		t.Run(test.description, func(t *testing.T) {
// 			cart := &core.Cart{}
// 			mockCouponService := mockNewCouponService(mockInput{applicableCoupon: "", discount: 0})
// 			cartService := New(cart, mockCouponService)

// 			cartService.AddProgram(test.input.quantity, test.input.category)
// 			received := cart.ProMembershipFee()
// 			if received != test.expected {
// 				t.Errorf("Expected %v, Received %v", test.expected, received)
// 			}
// 		})
// 	}

// }

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

func equal(t *testing.T, a, b []string) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
