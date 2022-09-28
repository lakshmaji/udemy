package printer_service

import (
	"bytes"
	writer_client "geektrust/clients/writer"
	"geektrust/core"
	"geektrust/core/coupon"
	"geektrust/core/program"
	"strings"
	"testing"
)

func TestBillTemplate(t *testing.T) {
	tt := []struct {
		description string
		input       *core.Cart
		expected    string
	}{
		{
			description: "when coupon applied",
			input: &core.Cart{
				Programs: []program.Program{
					{
						Category: program.CategoryDiploma,
						Quantity: 2,
					},
				},
				CouponApplied:         coupon.CouponDealG5,
				CouponDiscountApplied: 0.05 * 5000,
			},
			expected: func() string {
				var builder strings.Builder
				builder.WriteString("SUB_TOTAL\t5000.00\n")
				builder.WriteString("COUPON_DISCOUNT\tDEAL_G5\t250.00\n")
				builder.WriteString("TOTAL_PRO_DISCOUNT\t0.00\n")
				builder.WriteString("PRO_MEMBERSHIP_FEE\t0.00\n")
				builder.WriteString("ENROLLMENT_FEE\t500.00\n")
				builder.WriteString("TOTAL\t5250.00\n")
				return builder.String()
			}(),
		},
		{
			description: "no coupon applied",
			input: &core.Cart{
				Programs: []program.Program{
					{
						Category: program.CategoryCertification,
						Quantity: 1,
					},
				},
			},
			expected: func() string {
				var builder strings.Builder
				builder.WriteString("SUB_TOTAL\t3000.00\n")
				builder.WriteString("COUPON_DISCOUNT\tNONE\t0.00\n")
				builder.WriteString("TOTAL_PRO_DISCOUNT\t0.00\n")
				builder.WriteString("PRO_MEMBERSHIP_FEE\t0.00\n")
				builder.WriteString("ENROLLMENT_FEE\t500.00\n")
				builder.WriteString("TOTAL\t3500.00\n")
				return builder.String()
			}(),
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			var mockWriter bytes.Buffer
			writer := writer_client.New(&mockWriter, writer_client.DefaultOptions)

			printer := New(writer)
			printer.BillTemplate(test.input)

			received := mockWriter.String()
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
