package handlers

import (
	"bufio"
	"errors"
	"geektrust/clients"
	"geektrust/domain"
	"geektrust/service"
	"os"
	"strings"
)

type Command string

const (
	CommandAddProgram       = "ADD_PROGRAMME"
	CommandAddProMembership = "ADD_PRO_MEMBERSHIP"
	CommandApplyCoupon      = "APPLY_COUPON"
	CommandPrintBill        = "PRINT_BILL"
)

func HandleCart(writer clients.BaseWriter) {
	// TODO: need read inputs service (from http or shell)
	file, err := readInput()
	if err != nil {
		writer.WriteError(err)
	}
	defer file.Close()

	cart := &domain.Cart{}
	couponService := service.NewCouponService()
	cartService := service.NewCartService(cart, couponService)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)
		// FIXME: this [0] and [1] with regex matching to create dictionary
		switch argList[0] {
		case CommandAddProgram:
			cartService.AddProgram(argList)
		case CommandAddProMembership:
			cartService.AddProMembership()
		case CommandApplyCoupon:
			cartService.AddCoupon(argList)
		case CommandPrintBill:
			cartService.PrintBill(writer)
		default:
			writer.WriteError("Unrecognized command provided")
		}
	}
}

// FIXME: `os` should be loosely coupled
func readInput() (*os.File, error) {
	cliArgs := os.Args[1:]
	if len(cliArgs) == 0 {
		return nil, errors.New("Please provide the input file path")
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("Error opening the input file")
	}
	return file, nil
}
