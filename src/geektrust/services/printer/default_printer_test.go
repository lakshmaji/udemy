package printer

import (
	"bytes"
	writer_client "geektrust/clients/writer"
	"geektrust/core"
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
			description: "Empty cart",
			input:       &core.Cart{},
			expected: func() string {
				var builder strings.Builder
				builder.WriteString("SUB_TOTAL\t500.00\n")
				builder.WriteString("DISCOUNT\tNONE\t0\n")
				builder.WriteString("TOTAL_PRO_DISCOUNT\t0.00\n")
				builder.WriteString("PRO_MEMBERSHIP_FEE\t0.00\n")
				builder.WriteString("ENROLLMENT_FEE\t500.00\n")
				builder.WriteString("TOTAL\t500.00\n")
				return builder.String()
			}(),
		},
		{
			description: "No coupon applied",
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
				builder.WriteString("SUB_TOTAL\t3500.00\n")
				builder.WriteString("DISCOUNT\tNONE\t0\n")
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
			var output bytes.Buffer
			var writer writer_client.BaseWriter = writer_client.New(&output, writer_client.DefaultOptions)

			printer := New(writer)
			printer.BillTemplate(test.input)

			received := output.String()
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
