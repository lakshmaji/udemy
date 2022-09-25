package coupon

import (
	"fmt"
	"geektrust/core/coupon"
	"testing"
)

type testInput struct {
	noOfPrograms int
	subTotal     float64
	coupons      []coupon.Coupon
}

func TestApplicableCoupon_WhenNoCouponApplied(t *testing.T) {

	tt := []struct {
		description string
		input       testInput
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs are equal to %d, it should auto apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: testInput{
				noOfPrograms: 4,
				subTotal:     3000,
			},
			expected: coupon.CouponB4G1,
		},
		{
			description: fmt.Sprintf("when total no of programs are more than %d, it should auto apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: testInput{
				noOfPrograms: 6,
				subTotal:     19000,
			},
			expected: coupon.CouponB4G1,
		},
		{
			description: fmt.Sprintf("when total no of programs are less than %d, it should not apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: testInput{
				noOfPrograms: 2,
				subTotal:     5000,
			},
			expected: coupon.Coupon(""),
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			applicableCouponHelper(t, test.input, test.expected)
		})
	}
}

func TestApplicableCoupon_WhenB1G4Applied(t *testing.T) {

	tt := []struct {
		description string
		input       testInput
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs are less than to %d, it should not apply %s", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: testInput{
				noOfPrograms: 2,
				subTotal:     3000,
				coupons: []coupon.Coupon{
					coupon.CouponB4G1,
				},
			},
			expected: coupon.Coupon(""),
		},
		{
			description: fmt.Sprintf("when total no of programs are more than to %d, it should apply %s", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: testInput{
				noOfPrograms: 5,
				subTotal:     13000,
				coupons: []coupon.Coupon{
					coupon.CouponB4G1,
				},
			},
			expected: coupon.CouponB4G1,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			applicableCouponHelper(t, test.input, test.expected)
		})
	}
}

func TestApplicableCoupon_WhenDEAL_G20Applied(t *testing.T) {

	tt := []struct {
		description string
		input       testInput
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when sub total amount is equal to %d", coupon.CouponDealG20MarginAmount),
			input: testInput{
				noOfPrograms: 2,
				subTotal:     10000,
				coupons: []coupon.Coupon{
					coupon.CouponDealG20,
				},
			},
			expected: coupon.CouponDealG20,
		},
		{
			description: fmt.Sprintf("when sub total amount is more than %d", coupon.CouponDealG20MarginAmount),
			input: testInput{
				noOfPrograms: 2,
				subTotal:     12000,
				coupons: []coupon.Coupon{
					coupon.CouponDealG20,
				},
			},
			expected: coupon.CouponDealG20,
		},
		{
			description: fmt.Sprintf("when sub total amount is less than %d", coupon.CouponDealG20MarginAmount),
			input: testInput{
				noOfPrograms: 2,
				subTotal:     9800,
				coupons: []coupon.Coupon{
					coupon.CouponDealG20,
				},
			},
			expected: coupon.Coupon(""),
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			applicableCouponHelper(t, test.input, test.expected)
		})
	}
}

func TestApplicableCoupon_WhenDEAL_G5Applied(t *testing.T) {

	tt := []struct {
		description string
		input       testInput
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs is equal to %d, it should apply coupon", coupon.CouponDealG5MarginCount),
			input: testInput{
				noOfPrograms: 2,
				subTotal:     10000,
				coupons: []coupon.Coupon{
					coupon.CouponDealG5,
				},
			},
			expected: coupon.CouponDealG5,
		},
		{
			description: fmt.Sprintf("when total no of programs is less than %d, it should not apply coupon", coupon.CouponDealG5MarginCount),
			input: testInput{
				noOfPrograms: 1,
				subTotal:     10000,
				coupons: []coupon.Coupon{
					coupon.CouponDealG5,
				},
			},
			expected: coupon.Coupon(""),
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			applicableCouponHelper(t, test.input, test.expected)
		})
	}
}

func applicableCouponHelper(t *testing.T, input testInput, expected coupon.Coupon) {
	t.Helper()
	printer := New()
	received := printer.ApplicableCoupon(input.noOfPrograms, input.subTotal, input.coupons)

	if received != expected {
		t.Errorf("Expected %v, Received %v", expected, received)
	}
}
