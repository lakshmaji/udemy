package handlers

import (
	"geektrust/clients"
	"geektrust/domain"
	"geektrust/service"
)

// This will handle in applying and integrating the required services
// to process the input commands and returns (STDOUT) the total amount payable.
func CartHandler(writer clients.BaseWriter, reader clients.BaseReader) {
	var commands [][]string
	var err error

	commands, err = reader.CartCommands()
	if err != nil {
		writer.WriteError(err)
	}

	cart := &domain.Cart{}
	couponService := service.NewCouponService()
	cartService := service.NewCartService(cart, couponService)

	for _, command := range commands {
		switch domain.Command(command[0]) {
		case domain.CommandAddProgram:
			cartService.AddProgram(command)
		case domain.CommandAddProMembership:
			cartService.AddProMembership()
		case domain.CommandApplyCoupon:
			cartService.AddCoupon(command)
		case domain.CommandPrintBill:
			cartService.ComputeDiscountAndPrintBill(writer)
		default:
			writer.WriteError("Unrecognized command provided")
		}
	}

}
