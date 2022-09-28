package cart_service

import (
	"geektrust/core"
	"geektrust/core/coupon"
	"geektrust/core/program"
	"geektrust/services/coupon_service"
	"geektrust/utils"
	"strconv"
)

type service struct {
	cart          *core.Cart
	couponService coupon_service.CouponService
}

// New - creates new cart service with cart and coupon service implementation references.
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
	discount := c.couponService.CalculateDiscount(coupon, c.cart.Programs, subTotal, c.cart.HasProMemberShip)
	c.cart.CouponDiscountApplied = discount
	c.cart.CouponApplied = coupon
}
