package printer

import "geektrust/domain"

type PrinterService interface {
	// Generates the printable bill template for programs, discount and totals.
	// Invoke writer to write to STDOUT.
	BillTemplate(cart *domain.Cart)
}
