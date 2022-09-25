package cart

import (
	"geektrust/core/coupon"
	"geektrust/core/program"
	coupon_service "geektrust/services/coupon"
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

func (f *mockClient) CalculateDiscount(code coupon.Coupon, programs []program.Program, subTotal float64) float64 {
	return f.discount
}
