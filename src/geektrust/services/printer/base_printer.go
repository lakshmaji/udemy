/*
printer package handles final response format for bill.
*/
package printer

import "geektrust/core"

type PrinterService interface {
	// Generates the printable bill template for programs, discount and totals.
	//
	// Prints the total billable breakdown in a template defined in the concrete implementation.
	//
	// @param cart - cart instance as parameter.
	BillTemplate(cart *core.Cart)
}
