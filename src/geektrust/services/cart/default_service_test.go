package cart

import (
	"geektrust/core"
	"geektrust/core/coupon"
	"reflect"
	"testing"
)

func TestAddCoupon(t *testing.T) {
	mockCouponService := mockNewCouponService(mockInput{applicableCoupon: "", discount: 0})
	cart := &core.Cart{}
	cartService := New(cart, mockCouponService)
	cartService.AddCoupon("DEAL_NEW_ACCOUNT")
	expected := coupon.Coupons{coupon.Coupon("DEAL_NEW_ACCOUNT")}
	if !reflect.DeepEqual(cart.CouponsList, expected) {
		t.Errorf("Expected %v, Received %v", expected, cart.CouponsList)
	}
}

func equal(t *testing.T, a, b []string) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
