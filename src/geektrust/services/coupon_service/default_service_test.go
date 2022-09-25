package coupon_service

import (
	"fmt"
	"geektrust/core/coupon"
	"geektrust/core/program"
	"testing"
)

type inputApplicableCoupon struct {
	noOfPrograms int
	subTotal     float64
	coupons      []coupon.Coupon
}

type inputCalculateDiscount struct {
	code     coupon.Coupon
	programs []program.Program
	subTotal float64
}

func TestApplicableCoupon_WhenNoCouponApplied(t *testing.T) {

	tt := []struct {
		description string
		input       inputApplicableCoupon
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs are equal to %d, it should auto apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: inputApplicableCoupon{
				noOfPrograms: 4,
				subTotal:     3000,
			},
			expected: coupon.CouponB4G1,
		},
		{
			description: fmt.Sprintf("when total no of programs are more than %d, it should auto apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: inputApplicableCoupon{
				noOfPrograms: 6,
				subTotal:     19000,
			},
			expected: coupon.CouponB4G1,
		},
		{
			description: fmt.Sprintf("when total no of programs are less than %d, it should not apply %s coupon", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: inputApplicableCoupon{
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
		input       inputApplicableCoupon
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs are less than to %d, it should not apply %s", coupon.CouponB4G1MarginCount, coupon.CouponB4G1),
			input: inputApplicableCoupon{
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
			input: inputApplicableCoupon{
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
		input       inputApplicableCoupon
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when sub total amount is equal to %f", coupon.CouponDealG20MarginAmount),
			input: inputApplicableCoupon{
				noOfPrograms: 2,
				subTotal:     10000,
				coupons: []coupon.Coupon{
					coupon.CouponDealG20,
				},
			},
			expected: coupon.CouponDealG20,
		},
		{
			description: fmt.Sprintf("when sub total amount is more than %f", coupon.CouponDealG20MarginAmount),
			input: inputApplicableCoupon{
				noOfPrograms: 2,
				subTotal:     12000,
				coupons: []coupon.Coupon{
					coupon.CouponDealG20,
				},
			},
			expected: coupon.CouponDealG20,
		},
		{
			description: fmt.Sprintf("when sub total amount is less than %f", coupon.CouponDealG20MarginAmount),
			input: inputApplicableCoupon{
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
		input       inputApplicableCoupon
		expected    coupon.Coupon
	}{
		{
			description: fmt.Sprintf("when total no of programs is equal to %d, it should apply coupon", coupon.CouponDealG5MarginCount),
			input: inputApplicableCoupon{
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
			input: inputApplicableCoupon{
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

func applicableCouponHelper(t *testing.T, input inputApplicableCoupon, expected coupon.Coupon) {
	t.Helper()
	printer := New()
	received := printer.ApplicableCoupon(input.noOfPrograms, input.subTotal, input.coupons)

	if received != expected {
		t.Errorf("Expected %v, Received %v", expected, received)
	}
}

func TestCalculateDiscount(t *testing.T) {

	tt := []struct {
		description string
		input       inputCalculateDiscount
		expected    float64
	}{
		{
			description: fmt.Sprintf("coupon code %s", coupon.CouponDealG20),
			input: inputCalculateDiscount{
				code:     coupon.CouponDealG20,
				subTotal: 3000,
			},
			expected: 600.00,
		},
		{
			description: fmt.Sprintf("coupon code %s", coupon.CouponDealG5),
			input: inputCalculateDiscount{
				code:     coupon.CouponDealG5,
				subTotal: 3000,
			},
			expected: 150.00,
		},
		{
			description: fmt.Sprintf("coupon code %s", coupon.CouponB4G1),
			input: inputCalculateDiscount{
				code:     coupon.CouponB4G1,
				subTotal: 3000,
			},
			expected: 0,
		},
		{
			description: fmt.Sprintf("coupon code %s with some programs with fee's specified", coupon.CouponB4G1),
			input: inputCalculateDiscount{
				code:     coupon.CouponB4G1,
				subTotal: 3000,
				programs: []program.Program{
					{
						Quantity: 2,
						Category: program.CategoryCertification,
					},
					{
						Quantity: 1,
						Category: program.CategoryDiploma,
					},
					{
						Quantity: 1,
						Category: program.CategoryDegree,
					},
				},
			},
			expected: 2500,
		},
		{
			description: fmt.Sprintf("coupon code %s", coupon.Coupon("DEAL_G30")),
			input: inputCalculateDiscount{
				code:     coupon.Coupon("DEAL_G30"),
				subTotal: 3000,
			},
			expected: 0,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			applicableCalculateDiscount(t, test.input, test.expected)
		})
	}
}

func applicableCalculateDiscount(t *testing.T, input inputCalculateDiscount, expected float64) {
	t.Helper()
	printer := New()
	received := printer.CalculateDiscount(input.code, input.programs, input.subTotal)

	if received != expected {
		t.Errorf("Expected %v, Received %v", expected, received)
	}
}
