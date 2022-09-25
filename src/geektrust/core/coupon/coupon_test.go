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
