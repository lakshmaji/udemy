package coupon_service

import (
	"geektrust/core/coupon"
	"geektrust/core/program"
	"math"
)

type service struct {
}

// New - creates coupon service. This service has methods for choosing which coupon is eligible for entire catalog and to compute coupon discount.
func New() CouponService {
	return &service{}
}

func (c *service) ApplicableCoupon(count int, subTotal float64, coupons coupon.Coupons) coupon.Coupon {
	if count >= coupon.CouponB4G1MarginCount {
		return coupon.CouponB4G1
	}

	maxDiscountCoupon := coupons.Pick()

	var couponApplicable coupon.Coupon
	if subTotal >= coupon.CouponDealG20MarginAmount && maxDiscountCoupon == coupon.CouponDealG20 {
		couponApplicable = coupon.CouponDealG20
	} else if count >= coupon.CouponDealG5MarginCount && maxDiscountCoupon == coupon.CouponDealG5 {
		couponApplicable = coupon.CouponDealG5
	}

	return couponApplicable
}

func (c *service) CalculateDiscount(code coupon.Coupon, programs []program.Program, subTotal float64) float64 {
	var discount float64
	switch code {
	case coupon.CouponB4G1:
		var programMinCost float64 = math.MaxFloat64
		for _, p := range programs {
			if p.Category.Fee() < programMinCost {
				programMinCost = p.Category.Fee()
			}
		}
		if programMinCost == math.MaxFloat64 {
			return 0
		}
		discount = programMinCost
	case coupon.CouponDealG20:
		discount = subTotal * code.Percentage()
	case coupon.CouponDealG5:
		discount = subTotal * code.Percentage()
	default:
		discount = 0
	}
	return discount
}
