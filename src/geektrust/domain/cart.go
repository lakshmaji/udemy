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

// =============================================================================
// 						Enrollment and pro-membership fee
// =============================================================================

func (c *Cart) EnrollmentFee() float64 {
	// If the total programme cost is or above Rs.6666/- the enrollment fee is waived off.
	if (c.totalCostForPrograms() + c.ProMembershipFee() - c.TotalProMembershipDiscount()) <= EnrollmentFeeMargin {
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

func (c *Cart) totalCostForPrograms() float64 {
	var total float64
	for _, p := range c.Programs {
		total += p.Category.Fee() * float64(p.Qty)
	}
	return total
}

func (c *Cart) TotalProMembershipDiscount() float64 {
	var total float64
	for _, p := range c.Programs {
		total += p.ProMembershipDiscount(c.hasProMemberShip)
	}
	return total
}

func (c *Cart) SubTotal() float64 {
	var subTotal float64
	// This includes all the items purchased (programmes and pro membership)
	subTotal += c.totalCostForPrograms()
	subTotal += c.ProMembershipFee()
	// Print the cost of all the programmes purchased, after applying a pro membership discount (if applicable).
	subTotal -= c.TotalProMembershipDiscount()
	subTotal += c.EnrollmentFee()
	return subTotal
}

func (c *Cart) Total() float64 {
	return c.SubTotal() - c.CouponDiscountApplied
}

func (c *Cart) TotalProgramsCount() int {
	var count int
	for _, p := range c.Programs {
		count += p.Qty
	}
	return count
}

// =============================================================================
// 								Coupon utils
// =============================================================================

// Applies the discount coupon to the total value of the purchases
// func (c *Cart) SetEligibleCoupon() {
// 	noOfPrograms := c.totalPrograms()
// 	if noOfPrograms >= 4 {
// 		c.EligibleCoupon = coupon.CouponB4G1
// 		return
// 	}

// 	subTotal := c.SubTotal()
// 	couponConsidered := c.couponsList.Max()

// 	// FIXME: magic numbers
// 	var couponApplicable coupon.Coupon
// 	if subTotal >= 10000 && couponConsidered == coupon.CouponDealG20 {
// 		couponApplicable = coupon.CouponDealG20
// 	} else if noOfPrograms >= 2 && couponConsidered == coupon.CouponDealG5 {
// 		couponApplicable = coupon.CouponDealG5
// 	}

// 	c.EligibleCoupon = couponApplicable
// }

// func (c *Cart) SetCouponDiscount() {
// 	subTotal := c.SubTotal()
// 	switch c.EligibleCoupon {
// 	case coupon.CouponB4G1:
// 		var lowVal float64 = math.MaxFloat64
// 		for _, p := range c.programs {
// 			if p.Category.Fee() < lowVal {
// 				lowVal = p.Category.Fee()
// 			}
// 		}
// 		c.CouponDiscountApplied = lowVal
// 	case coupon.CouponDealG20:
// 		c.CouponDiscountApplied = subTotal * c.EligibleCoupon.Percentage()
// 	case coupon.CouponDealG5:
// 		c.CouponDiscountApplied = subTotal * c.EligibleCoupon.Percentage()
// 	default:
// 		c.CouponDiscountApplied = 0
// 	}
// }
