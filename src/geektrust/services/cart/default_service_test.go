package cart

import (
	"geektrust/core"
	"geektrust/core/coupon"
	"reflect"
	"testing"
)

func TestAddCoupon(t *testing.T) {

	tt := []struct {
		description        string
		input              []string
		couponSvcMockInput mockInput
		expected           coupon.Coupons
	}{
		{
			description:        "invalid coupon added",
			input:              []string{"DEAL_NEW_ACCOUNT"},
			couponSvcMockInput: mockInput{applicableCoupon: "", discount: 0},
			expected:           coupon.Coupons{coupon.Coupon("DEAL_NEW_ACCOUNT")},
		},
		{
			description:        "valid coupon added",
			input:              []string{"DEAL_G20"},
			couponSvcMockInput: mockInput{applicableCoupon: "", discount: 0},
			expected:           coupon.Coupons{coupon.Coupon("DEAL_G20")},
		},
		{
			description:        "add multiple coupons",
			input:              []string{"DEAL_G20", "B4G1"},
			couponSvcMockInput: mockInput{applicableCoupon: "", discount: 0},
			expected:           coupon.Coupons{coupon.Coupon("DEAL_G20"), coupon.Coupon("B4G1")},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			cart := &core.Cart{}
			mockCouponService := mockNewCouponService(test.couponSvcMockInput)
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
