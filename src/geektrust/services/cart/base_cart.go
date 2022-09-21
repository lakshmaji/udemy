package cart

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
