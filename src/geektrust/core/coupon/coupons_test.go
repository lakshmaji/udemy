package coupon

import (
	"testing"
)

func TestPick(t *testing.T) {
	coupons := Coupons{CouponDealG20}

	received := coupons.Pick()
	expected := CouponDealG20
	t.Log(received)
	if expected != received {
		t.Errorf("Expected %v, Received %v", expected, received.Percentage())
	}
}
