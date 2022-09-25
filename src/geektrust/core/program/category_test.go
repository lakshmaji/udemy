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
