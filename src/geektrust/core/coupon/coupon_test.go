package coupon

import "testing"

func TestCouponDealG20(t *testing.T) {
	expected := Coupon("DEAL_G20")
	if CouponDealG20 != expected {
		t.Errorf("Expected %s, Received %s", expected, CouponDealG20)
	}
}
func TestCouponDealG5(t *testing.T) {
	expected := Coupon("DEAL_G5")
	if CouponDealG5 != expected {
		t.Errorf("Expected %s, Received %s", expected, CouponDealG5)
	}
}
func TestCouponB4G1(t *testing.T) {
	expected := Coupon("B4G1")
	if CouponB4G1 != expected {
		t.Errorf("Expected %s, Received %s", expected, CouponB4G1)
	}
}

func TestCouponDealG20MarginAmount(t *testing.T) {
	expected := float64(10000)
	if CouponDealG20MarginAmount != expected {
		t.Errorf("Expected %f, Received %f", expected, CouponDealG20MarginAmount)
	}
}

func TestCouponDealG5MarginCount(t *testing.T) {
	expected := 2
	if CouponDealG5MarginCount != expected {
		t.Errorf("Expected %d, Received %d", expected, CouponDealG5MarginCount)
	}
}

func TestCouponB4G1MarginCount(t *testing.T) {
	expected := 4
	if CouponB4G1MarginCount != expected {
		t.Errorf("Expected %d, Received %d", expected, CouponB4G1MarginCount)
	}
}

func TestDealG20(t *testing.T) {
	expected := 20 / float64(100)
	if dealG20 != expected {
		t.Errorf("Expected %f, Received %f", expected, dealG20)
	}
}

func TestDealG5(t *testing.T) {
	expected := 5 / float64(100)
	if dealG5 != expected {
		t.Errorf("Expected %f, Received %f", expected, dealG5)
	}
}

func TestPercentage(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    float64
	}{
		{
			description: "DEAL_G20",
			input:       "DEAL_G20",
			expected:    0.2,
		},
		{
			description: "DEAL_G5",
			input:       "DEAL_G5",
			expected:    0.05,
		},
		{
			description: "B4G1",
			input:       "B4G1",
			expected:    0,
		},
		{
			description: "UNKNOWN",
			input:       "UNKNOWN",
			expected:    0,
		},
	}
	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			received := Coupon(test.input).Percentage()
			if received != test.expected {
				t.Errorf("Expected %v, Received %v", test.expected, received)
			}
		})
	}
}
