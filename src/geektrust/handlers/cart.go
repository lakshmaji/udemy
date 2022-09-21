package handlers

import (
	reader_client "geektrust/clients/reader"
	writer_client "geektrust/clients/writer"
	"geektrust/domain"
	cart_service "geektrust/services/cart"
	"geektrust/services/command_parser"
	coupon_service "geektrust/services/coupon"
	printer_service "geektrust/services/printer"
)

// This will handle in applying and integrating the required services
// to process the input commands and returns (STDOUT) the total amount payable.
func CartHandler(writer writer_client.BaseWriter, reader reader_client.BaseReader) {
	cart := &domain.Cart{}

	// Initialize services
	couponService := coupon_service.New()
	cartService := cart_service.New(cart, couponService)
	commandIOService := command_parser.New(reader)

	// Get input commands
	commands, err := commandIOService.Commands()
	if err != nil {
		writer.WriteError(err)
	}

	// Process commands
	for _, command := range commands {
		switch domain.Command(command[0]) {
		case domain.CommandAddProgram:
			cartService.AddProgram(command)
		case domain.CommandAddProMembership:
			cartService.AddProMembership()
		case domain.CommandApplyCoupon:
			cartService.AddCoupon(command)
		case domain.CommandPrintBill:
			cartService.ComputeDiscount()
			printer := printer_service.NewPrinterService(writer)
			printer.BillTemplate(cart)
		default:
			writer.WriteError("Unrecognized command provided")
		}
	}

}
