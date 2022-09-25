package utils

import (
	"geektrust/core/program"
	"testing"
)

func TestMapStringToProgramCategory(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    program.Category
	}{
		{
			description: "Certification",
			input:       "CERTIFICATION",
			expected:    program.CategoryCertification,
		},
		{
			description: "Degree",
			input:       "DEGREE",
			expected:    program.CategoryDegree,
		},
		{
			description: "Diploma",
			input:       "DIPLOMA",
			expected:    program.CategoryDiploma,
		},
		{
			description: "Invalid",
			input:       "COMPUTERS",
			expected:    program.CategoryUnknown,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := MapStringToProgramCategory(test.input)

			if received != test.expected {
				t.Errorf("Expected %s, Received %s", test.expected, received)
			}
		})
	}
}
