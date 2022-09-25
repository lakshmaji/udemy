package program

import "testing"

// Pro membership discounts
func TestProDiscountDiploma(t *testing.T) {
	if ProDiscountDiploma != 0.01 {
		t.Error("should be ", ProDiscountDiploma)
	}
}
func TestProDiscountCertification(t *testing.T) {
	if ProDiscountCertification != 0.02 {
		t.Error("should be ", ProDiscountCertification)
	}
}
func TestProDiscountDegree(t *testing.T) {
	if ProDiscountDegree != 0.03 {
		t.Error("should be ", ProDiscountDegree)
	}
}

// category amounts
func TestCostCertification(t *testing.T) {
	if CostCertification != 3000 {
		t.Error("Should be 3000")
	}
}
func TestCostDegree(t *testing.T) {
	if CostDegree != 5000 {
		t.Error("Should be 5000")
	}
}
func TestCostDiploma(t *testing.T) {
	if CostDiploma != 2500 {
		t.Error("Should be 2500")
	}
}

func TestCategoryString(t *testing.T) {
	tt := []struct {
		description string
		input       Category
		expected    string
	}{
		{
			description: "certification",
			input:       CategoryCertification,
			expected:    "CERTIFICATION",
		},
		{
			description: "degree",
			input:       CategoryDegree,
			expected:    "DEGREE",
		},
		{
			description: "diploma",
			input:       CategoryDiploma,
			expected:    "DIPLOMA",
		},
		{
			description: "invalid",
			input:       CategoryUnknown,
			expected:    "unknown category",
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := test.input.String()
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}

func TestCategoryFee(t *testing.T) {
	tt := []struct {
		description string
		input       Category
		expected    float64
	}{
		{
			description: "certification",
			input:       CategoryCertification,
			expected:    3000,
		},
		{
			description: "degree",
			input:       CategoryDegree,
			expected:    5000,
		},
		{
			description: "diploma",
			input:       CategoryDiploma,
			expected:    2500,
		},
		{
			description: "invalid",
			input:       CategoryUnknown,
			expected:    0,
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := test.input.Fee()
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
