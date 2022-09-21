package service

import (
	"geektrust/domain/coupon"
	"geektrust/domain/program"
	"math"
)

type CouponService interface {

	// Selects the applicable coupon on the cart items (programs)
	// @param count 	- no programs in the cart
	// @param subTotal 	- the subTotal for your cart
	// @param coupons 	- the list of coupons applied on the cart
	ApplicableCoupon(count int, subTotal float64, coupons coupon.Coupons) coupon.Coupon

	// Calculate the total discount for a applied coupon code on the provided amount
	// @param amount 	- Could be subTotal
	// @param programs 	- the program items in the cart
	// @param code 		- Coupon code to apply
	CalculateDiscount(code coupon.Coupon, programs []program.Program, amount float64) float64
}

type couponSvc struct {
}

func NewCouponService() CouponService {
	return &couponSvc{}
}

// Selects the applicable coupon on the cart items (programs)
// @param count 	- no programs in the cart
// @param subTotal 	- the subTotal for your cart
// @param coupons 	- the list of coupons applied on the cart
func (c *couponSvc) ApplicableCoupon(count int, subTotal float64, coupons coupon.Coupons) coupon.Coupon {
	if count >= coupon.CouponB4G1MarginCount {
		return coupon.CouponB4G1
	}

	maxDiscountCoupon := coupons.Max()

	var couponApplicable coupon.Coupon
	if subTotal >= coupon.CouponDealG20MarginAmount && maxDiscountCoupon == coupon.CouponDealG20 {
		couponApplicable = coupon.CouponDealG20
	} else if count >= coupon.CouponDealG5MarginCount && maxDiscountCoupon == coupon.CouponDealG5 {
		couponApplicable = coupon.CouponDealG5
	}

	return couponApplicable
}

// Calculate the total discount for a applied coupon code on the provided amount
// @param amount 	- Could be subTotal
// @param programs 	- the program items in the cart
// @param code 		- Coupon code to apply
func (c *couponSvc) CalculateDiscount(code coupon.Coupon, programs []program.Program, amount float64) float64 {
	var discount float64
	switch code {
	case coupon.CouponB4G1:
		var programMinCost float64 = math.MaxFloat64
		for _, p := range programs {
			if p.Category.Fee() < programMinCost {
				programMinCost = p.Category.Fee()
			}
		}
		discount = programMinCost
	case coupon.CouponDealG20:
		discount = amount * code.Percentage()
	case coupon.CouponDealG5:
		discount = amount * code.Percentage()
	default:
		discount = 0
	}
	return discount
}
