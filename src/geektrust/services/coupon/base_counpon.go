/*
This coupon package handle discount applying logic and selecting the final eligible coupon for the cart.
*/
package coupon

import (
	"geektrust/core/coupon"
	"geektrust/core/program"
)

type CouponService interface {

	// Returns the applicable coupon on the cart items (programs)
	//
	// @param count 	- no programs in the cart
	//
	// @param subTotal 	- the subTotal for your cart
	//
	// @param coupons 	- the list of coupons applied on the cart
	ApplicableCoupon(count int, subTotal float64, coupons coupon.Coupons) coupon.Coupon

	// Returns the calculated total discount for a given coupon code on the provided amount
	//
	// @param subTotal 	- sub total (net amount)
	//
	// @param programs 	- the program items in the cart
	//
	// @param code 		- Coupon code to apply
	CalculateDiscount(code coupon.Coupon, programs []program.Program, subTotal float64) float64
}
