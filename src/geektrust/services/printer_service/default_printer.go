package printer_service

import (
	writer_client "geektrust/clients/writer"
	"geektrust/core"
)

// This expects a BaseWriter implementation, which will be used to write the output.
type printer struct {
	writer writer_client.BaseWriter
}

// New - This creates a printer service instance. A BaseWrite implementation should be passed in the argument.
func New(writer writer_client.BaseWriter) Printer {
	return &printer{writer}
}

func (p *printer) BillTemplate(cart *core.Cart) {
	p.writer.WriteLn("SUB_TOTAL\t%.2f", cart.SubTotal())
	if cart.CouponApplied == "" {
		p.writer.WriteLn("COUPON_DISCOUNT\tNONE\t0.00")
	} else {
		p.writer.WriteLn("COUPON_DISCOUNT\t%s\t%.2f", cart.CouponApplied, cart.CouponDiscountApplied)
	}
	p.writer.WriteLn("TOTAL_PRO_DISCOUNT\t%.2f", cart.TotalProMembershipDiscount())
	p.writer.WriteLn("PRO_MEMBERSHIP_FEE\t%.2f", cart.ProMembershipFee())
	p.writer.WriteLn("ENROLLMENT_FEE\t%.2f", cart.EnrollmentFee())
	p.writer.WriteLn("TOTAL\t%.2f", cart.NetTotal())
}
