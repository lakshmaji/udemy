/*
Package printersvc - package handles final response format for bill.
*/
package printer_service

import "geektrust/core"

// Printer ...
type Printer interface {
	// Generates the printable bill template for programs, discount and totals.
	//
	// Prints the total billable breakdown in a template defined in the concrete implementation.
	//
	// @param cart - cart instance as parameter.
	BillTemplate(cart *core.Cart)
}
