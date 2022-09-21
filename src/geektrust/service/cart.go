package service

import (
	"geektrust/clients"
	"geektrust/domain"
	"geektrust/domain/coupon"
	"geektrust/domain/program"
	"geektrust/printer"
	"geektrust/utils"
	"log"
	"strconv"
)

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
	ComputeDiscountAndPrintBill(writer clients.BaseWriter)
}

type cartSvc struct {
	cart          *domain.Cart
	couponService CouponService
}

func NewCartService(cart *domain.Cart, couponService CouponService) CartService {
	return &cartSvc{cart, couponService}
}

// add coupon
func (c *cartSvc) AddCoupon(argList []string) {
	c.cart.AddCoupon(coupon.Coupon(argList[1]))
}

// Add pro membership
func (c *cartSvc) AddProMembership() {
	c.cart.AddProMembership()
}

// Add item (program) to the cart.
func (c *cartSvc) AddProgram(argList []string) {
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
func (c *cartSvc) ComputeDiscountAndPrintBill(writer clients.BaseWriter) {
	// Apply discount and computes the total, before printing.
	subTotal := c.cart.SubTotal()
	coupon := c.couponService.ApplicableCoupon(c.cart.TotalProgramsCount(), subTotal, c.cart.CouponsList)
	discount := c.couponService.CalculateDiscount(coupon, c.cart.Programs, subTotal)
	c.cart.CouponDiscountApplied = discount
	c.cart.CouponApplied = coupon

	print := printer.NewPrinter(writer)
	print.BillTemplate(c.cart)
}
