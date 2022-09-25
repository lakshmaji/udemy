package program

import "testing"

func TestProMembershipDiscount(t *testing.T) {
	tt := []struct {
		description   string
		input         Program
		hasMembership bool
		expected      float64
	}{
		{
			description: "no pro-membership",
			input: Program{
				Category: CategoryCertification,
				Quantity: 2,
			},
			hasMembership: false,
			expected:      0,
		},
		{
			description: "has pro-membership",
			input: Program{
				Category: CategoryCertification,
				Quantity: 2,
			},
			hasMembership: true,
			expected:      120,
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := test.input.ProMembershipDiscount(test.hasMembership)
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
