package coupon

import (
	"fmt"
	"geektrust/core/coupon"
	"testing"
)

func TestApplicableCoupon(t *testing.T) {
	type TestInput struct {
		noOfPrograms int
		subTotal     float64
		coupons      []coupon.Coupon
	}

	tt := []struct {
		description string
		input       TestInput
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs are equal to %d, it should auto apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: TestInput{
				noOfPrograms: 4,
				subTotal:     3000,
			},
			expected: coupon.CouponB4G1,
		},
		{
			description: fmt.Sprintf("when total no of programs are more than %d, it should auto apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: TestInput{
				noOfPrograms: 6,
				subTotal:     19000,
			},
			expected: coupon.CouponB4G1,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {

			printer := New()
			received := printer.ApplicableCoupon(test.input.noOfPrograms, test.input.subTotal, test.input.coupons)

			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
