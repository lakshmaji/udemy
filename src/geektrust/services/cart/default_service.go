package cart

import (
	"geektrust/core"
	"geektrust/core/coupon"
	"geektrust/core/program"
	coupon_service "geektrust/services/coupon"
	"geektrust/utils"
	"strconv"
)

type service struct {
	cart          *core.Cart
	couponService coupon_service.CouponService
}

func New(cart *core.Cart, couponService coupon_service.CouponService) CartService {
	return &service{cart, couponService}
}

func (c *service) AddCoupon(codeCmd string) {
	c.cart.AddCoupon(coupon.Coupon(codeCmd))
}

func (c *service) AddProMembership() {
	c.cart.AddProMembership()
}

func (c *service) AddProgram(quantityCmd string, categoryCmd string) error {
	qty, err := strconv.Atoi(quantityCmd)
	if err != nil {
		return err
	}

	category := utils.MapStringToProgramCategory(categoryCmd)
	p := program.Program{
		Category: category,
		Quantity: qty,
	}

	if err := c.cart.AddProgram(p); err != nil {
		return err
	}
	return nil
}

func (c *service) ComputeDiscount() {
	subTotal := c.cart.SubTotal()
	coupon := c.couponService.ApplicableCoupon(c.cart.TotalProgramsCount(), subTotal, c.cart.CouponsList)
	discount := c.couponService.CalculateDiscount(coupon, c.cart.Programs, subTotal)
	c.cart.CouponDiscountApplied = discount
	c.cart.CouponApplied = coupon
}
