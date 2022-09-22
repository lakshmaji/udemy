package core

// Command - The program command can be one of following
//
//	ADD_PROGRAMME
//	APPLY_COUPON
//	ADD_PRO_MEMBERSHIP
//	PRINT_BILL
type Command string

// Enum for program commands
const (
	CommandAddProgram       = "ADD_PROGRAMME"
	CommandAddProMembership = "ADD_PRO_MEMBERSHIP"
	CommandApplyCoupon      = "APPLY_COUPON"
	CommandPrintBill        = "PRINT_BILL"
)
