package domain

import (
	"errors"
	"geektrust/domain/coupon"
	"geektrust/domain/program"
)

const (
	EnrollmentFeeMargin float64 = 6666
	EnrollmentFee       float64 = 500
	ProMemberShipFee    float64 = 200
)

type Cart struct {
	Programs              []program.Program
	hasProMemberShip      bool
	CouponsList           coupon.Coupons
	CouponDiscountApplied float64
	CouponApplied         coupon.Coupon
}

// =============================================================================
// 						Enrollment and pro-membership fee
// =============================================================================

// If the total programme cost is or above Rs.6666/- the enrollment fee is waived off.
func (c *Cart) EnrollmentFee() float64 {
	if c.programsNetAmount() < EnrollmentFeeMargin {
		return EnrollmentFee
	}
	return 0
}

func (c *Cart) ProMembershipFee() float64 {
	if c.hasProMemberShip {
		return ProMemberShipFee
	}
	return 0
}

// =============================================================================
// 								compute utils
// =============================================================================

func (c *Cart) programsNetAmount() float64 {
	programsCost := c.programsGrossAmount()
	memberShipFee := c.ProMembershipFee()
	discount := c.TotalProMembershipDiscount()
	return programsCost + memberShipFee - discount
}

func (c *Cart) programsGrossAmount() float64 {
	var total float64
	for _, p := range c.Programs {
		total += p.Category.Fee() * float64(p.Qty)
	}
	return total
}

func (c *Cart) TotalProMembershipDiscount() float64 {
	var discount float64
	for _, p := range c.Programs {
		discount += p.ProMembershipDiscount(c.hasProMemberShip)
	}
	return discount
}

func (c *Cart) SubTotal() float64 {
	var total float64
	// This includes all the items purchased (programmes and pro membership)
	total += c.programsGrossAmount()
	total += c.ProMembershipFee()
	// Print the cost of all the programmes purchased, after applying a pro membership discount (if applicable).
	total -= c.TotalProMembershipDiscount()
	total += c.EnrollmentFee()
	return total
}

func (c *Cart) Total() float64 {
	return c.SubTotal() - c.CouponDiscountApplied
}

func (c *Cart) TotalProgramsCount() int {
	var noOfItems int
	for _, p := range c.Programs {
		noOfItems += p.Qty
	}
	return noOfItems
}

// =============================================================================
// 						Add program, coupon and membership
// =============================================================================

func (c *Cart) AddProMembership() {
	c.hasProMemberShip = true
}

func (c *Cart) AddProgram(p program.Program) error {
	if p.Category == program.CategoryUnknown {
		return errors.New("Unknown program category")
	}
	c.Programs = append(c.Programs, p)
	return nil
}

func (c *Cart) AddCoupon(code coupon.Coupon) {
	c.CouponsList = append(c.CouponsList, code)
}
