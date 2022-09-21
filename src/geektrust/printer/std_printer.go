package printer

import (
	"geektrust/clients"
	"geektrust/domain"
)

type Printer struct {
	writer clients.BaseWriter
}

func NewPrinter(writer clients.BaseWriter) PrinterFactory {
	return &Printer{writer}
}

func (p *Printer) PrintBill(cart *domain.Cart) {
	p.writer.WriteLn("SUB_TOTAL\t%.2f", cart.SubTotal())
	if cart.CouponApplied == "" {
		p.writer.WriteLn("DISCOUNT\tNONE\t0")
	} else {
		p.writer.WriteLn("COUPON_DISCOUNT\t%s\t%.2f", cart.CouponApplied, cart.CouponDiscountApplied)
	}
	p.writer.WriteLn("TOTAL_PRO_DISCOUNT\t%.2f", cart.TotalProMembershipDiscount())
	p.writer.WriteLn("PRO_MEMBERSHIP_FEE\t%.2f", cart.ProMembershipFee())
	p.writer.WriteLn("ENROLLMENT_FEE\t%.2f", cart.EnrollmentFee())
	p.writer.WriteLn("TOTAL\t%.2f", cart.Total())
}
