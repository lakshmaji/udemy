package cart

import (
	"geektrust/domain"
	"geektrust/domain/coupon"
	"geektrust/domain/program"
	coupon_service "geektrust/services/coupon"
	"geektrust/utils"
	"log"
	"strconv"
)

type service struct {
	cart          *domain.Cart
	couponService coupon_service.CouponService
}

func New(cart *domain.Cart, couponService coupon_service.CouponService) CartService {
	return &service{cart, couponService}
}

// add coupon
func (c *service) AddCoupon(argList []string) {
	c.cart.AddCoupon(coupon.Coupon(argList[1]))
}

// Add pro membership
func (c *service) AddProMembership() {
	c.cart.AddProMembership()
}

// Add item (program) to the cart.
func (c *service) AddProgram(argList []string) {
	qty, err := strconv.Atoi(argList[2])
	if err != nil {
		log.Fatal(err)
	}

	category := utils.MapStringToProgramCategory(argList[1])
	p := program.Program{
		Category: category,
		Qty:      qty,
	}

	if err := c.cart.AddProgram(p); err != nil {
		log.Fatal(err)
	}
}

// Computes discount using coupon applied (or considered)
// Prints the total billable breakdown in a template defined in `printer.PrintBill()`.
// This computes the coupon discount to be applicable (based on coupon eligibility criteria).
// Updates cart instance with discount applied on overall cart and coupon considered.
func (c *service) ComputeDiscount() {
	// Apply discount and computes the total, before printing.
	subTotal := c.cart.SubTotal()
	coupon := c.couponService.ApplicableCoupon(c.cart.TotalProgramsCount(), subTotal, c.cart.CouponsList)
	discount := c.couponService.CalculateDiscount(coupon, c.cart.Programs, subTotal)
	c.cart.CouponDiscountApplied = discount
	c.cart.CouponApplied = coupon
}
