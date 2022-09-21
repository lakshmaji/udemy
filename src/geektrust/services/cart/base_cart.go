package cart

type CartService interface {
	// Add item (program) to the cart.
	AddProgram(argList []string)
	// Add pro membership
	AddProMembership()
	// add coupon
	AddCoupon(argList []string)
	// Computes discount using coupon applied (or considered)
	// Prints the total billable breakdown in a template defined in `printer.PrintBill()`.
	// This computes the coupon discount to be applicable (based on coupon eligibility criteria).
	// Updates cart instance with discount applied on overall cart and coupon considered.
	ComputeDiscount()
}
