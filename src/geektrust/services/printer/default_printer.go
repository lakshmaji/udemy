package printer

import (
	writer_client "geektrust/clients/writer"
	"geektrust/domain"
)

// This expects a BaseWriter implementation, which will be used to write the output.
type printer struct {
	writer writer_client.BaseWriter
}

// This instantiate a printer service instance.
//
// Expects a BaseWrite implementation as function parameter.
func New(writer writer_client.BaseWriter) PrinterService {
	return &printer{writer}
}

func (p *printer) BillTemplate(cart *domain.Cart) {
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
