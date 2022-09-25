package core

import (
	"geektrust/core/coupon"
	"geektrust/core/program"
	"geektrust/utils"
)

const (
	// EnrollmentFeeMargin is a marginal (cart) amount for applying enrollment fee.
	EnrollmentFeeMargin float64 = 6666
	// EnrollmentFee indicates amount
	EnrollmentFee float64 = 500
	// ProMemberShipFee is the amount which is applied when student is subscribed to pro-membership.
	ProMemberShipFee float64 = 200
)

// Cart - holds and encapsulate the data with help of several members defined over it.
type Cart struct {
	Programs []program.Program
	// If students has added pro membership to cart
	hasProMemberShip bool
	// List of coupons applied by student
	CouponsList coupon.Coupons
	// Coupon discount after considering single coupon based on criteria provided.
	CouponDiscountApplied float64
	// Coupon selected based on criteria provided.
	CouponApplied coupon.Coupon
}

// EnrollmentFee - If the total programme cost is or above Rs.6666/- the enrollment fee is waived off.
func (c *Cart) EnrollmentFee() float64 {
	if c.programsNetAmount() < EnrollmentFeeMargin {
		return EnrollmentFee
	}
	return 0
}

// ProMembershipFee - a Pro Membership for a small amount of Rs.200/- is applicable when student purchase Pro Membership.
func (c *Cart) ProMembershipFee() float64 {
	if c.hasProMemberShip {
		return ProMemberShipFee
	}
	return 0
}

// Computes total netAmount for programs, membership and membership discount.
//
// netAmount = grossAmount (all programs) + pro membership fee - pro membership discount (all programs)
func (c *Cart) programsNetAmount() float64 {
	var total float64
	total += c.programsGrossAmount()
	total += c.ProMembershipFee()
	total -= c.TotalProMembershipDiscount()
	return total
}

// Computes gross amount for all the programmes category in the cart.
//
// grossAmount = program fee * quantity (all programs)
func (c *Cart) programsGrossAmount() float64 {
	var total float64
	for _, p := range c.Programs {
		total += p.Category.Fee() * float64(p.Quantity)
	}
	return total
}

// TotalProMembershipDiscount - Computes total membership discount if applicable.
func (c *Cart) TotalProMembershipDiscount() float64 {
	var discount float64
	for _, p := range c.Programs {
		discount += p.ProMembershipDiscount(c.hasProMemberShip)
	}
	return discount
}

// SubTotal - Computes subTotal
//
// total = program gross amount + membership fee - membership discount + enrollment fee
func (c *Cart) SubTotal() float64 {
	var total float64
	// This includes all the items purchased (programmes and pro membership)
	total += c.programsGrossAmount()
	total += c.ProMembershipFee()
	// Print the cost of all the programmes purchased, after applying a pro membership discount (if applicable).
	total -= c.TotalProMembershipDiscount()
	return total
}

// Total - Computes final payable amount
//
// total = subTotal - coupon discount
func (c *Cart) Total() float64 {
	return c.SubTotal() - c.CouponDiscountApplied + c.EnrollmentFee()
}

// TotalProgramsCount - The total no of programs in the cart.
func (c *Cart) TotalProgramsCount() int {
	var noOfItems int
	for _, p := range c.Programs {
		noOfItems += p.Quantity
	}
	return noOfItems
}

// AddProMembership - subscribe student to membership
func (c *Cart) AddProMembership() {
	c.hasProMemberShip = true
}

// AddProgram - add program category to cart with given quantity.
func (c *Cart) AddProgram(p program.Program) error {
	if p.Category == program.CategoryUnknown {
		return utils.ErrorUnknownCategory
	}
	c.Programs = append(c.Programs, p)
	return nil
}

// AddCoupon - apply a single coupon (at a time) to the entire cart.
// This coupon is used to deduct certain amount (which is based on coupon applied) from final amount (overall cart).
func (c *Cart) AddCoupon(code coupon.Coupon) {
	c.CouponsList = append(c.CouponsList, code)
}
