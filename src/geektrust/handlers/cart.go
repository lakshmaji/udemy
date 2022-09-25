package handlers

import (
	reader_client "geektrust/clients/reader"
	writer_client "geektrust/clients/writer"
	"geektrust/core"
	"geektrust/services/cart_service"
	"geektrust/services/cmd_service"
	"geektrust/services/coupon_service"
	"geektrust/utils"
)

// CartHandler - This will guide the program in applying and integrating the required services
// to process the input commands and returns (STDOUT) the total amount payable.
func CartHandler(writer writer_client.BaseWriter, reader reader_client.BaseReader) {
	cart := &core.Cart{}

	// Initialize services
	couponService := coupon_service.New()
	cartService := cart_service.New(cart, couponService)
	inputParser := cmd_service.New(reader)
	printer := printersvc.New(writer)

	// Get input commands
	commands, err := inputParser.Commands()
	if err != nil {
		writer.WriteError(err)
	}

	// Process commands
	for _, command := range commands {
		switch core.Command(command[0]) {
		case core.CommandAddProgram:
			err := cartService.AddProgram(command[2], command[1])
			if err != nil {
				writer.WriteError(err)
			}
		case core.CommandAddProMembership:
			cartService.AddProMembership()
		case core.CommandApplyCoupon:
			cartService.AddCoupon(command[1])
		case core.CommandPrintBill:
			cartService.ComputeDiscount()
			printer.BillTemplate(cart)
		default:
			writer.WriteError(&utils.UnknownCommandError{Command: command[0]})
		}
	}

}
