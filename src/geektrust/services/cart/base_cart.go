/*
Package cart - is a service, which is used to manipulate the cart instance properties.

Add program category to cart with specified number of items.

Add pro membership subscription.

Apply coupon code.

Computes discount after applying a eligible coupon.
*/
package cart

// CartService ...
type CartService interface {
	// Add item (program) to the cart.
	AddProgram(quantityCmd string, categoryCmd string) error
	// Add pro membership
	AddProMembership()
	// add coupon
	AddCoupon(codeCmd string)
	// Computes discount using coupon applied (or considered)
	//
	// This computes the coupon discount to be applicable (based on coupon eligibility criteria).
	//
	// Updates cart instance with discount applied on overall cart and coupon considered.
	ComputeDiscount()
}
