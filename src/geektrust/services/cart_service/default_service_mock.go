package cart_service

import (
	"geektrust/core/coupon"
	"geektrust/core/program"
	"geektrust/services/coupon_service"
)

type mockInput struct {
	applicableCoupon string
	discount         float64
}

type mockClient struct {
	mockInput
}

func mockNewCouponService(
	mockInput mockInput,
) coupon_service.CouponService {
	return &mockClient{mockInput: mockInput}
}

func (f *mockClient) ApplicableCoupon(count int, subTotal float64, coupons coupon.Coupons) coupon.Coupon {
	return coupon.Coupon(f.applicableCoupon)
}

func (f *mockClient) CalculateDiscount(code coupon.Coupon, programs []program.Program, subTotal float64, hasProMemberShip bool) float64 {
	return f.discount
}
