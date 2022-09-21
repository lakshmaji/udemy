package printer

import "geektrust/domain"

type PrinterFactory interface {
	// Writes output to provided writer client
	PrintBill(cart *domain.Cart)
}
