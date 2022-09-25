package coupon

import (
	"testing"
)

func TestPick(t *testing.T) {
	tt := []struct {
		description string
		coupons     Coupons
		expected    Coupon
	}{
		{
			description: "returns DEAL_G20 when DEAL_G20 is available ",
			coupons:     Coupons{CouponDealG20},
			expected:    CouponDealG20,
		},
		{
			description: "returns DEAL_G5 when DEAL_G5 is available",
			coupons:     Coupons{CouponDealG5},
			expected:    CouponDealG5,
		},
		{
			description: "returns B4G1 when B4G1 is available",
			coupons:     Coupons{CouponB4G1},
			expected:    CouponB4G1,
		},
		{
			description: "returns empty string when no coupon is available ",
			coupons:     Coupons{},
			expected:    Coupon(""),
		},
		{
			description: "returns max amount coupon when more coupons are available",
			coupons:     Coupons{CouponDealG5, CouponDealG20},
			expected:    CouponDealG20,
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := test.coupons.Pick()
			if test.expected != received {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}

}
