package core

import "testing"

func TestCommandAddProgram(t *testing.T) {
	received := CommandAddProgram
	expected := "ADD_PROGRAMME"
	if expected != received {
		t.Errorf("Expected %s, Received %s", expected, received)
	}
}
func TestCommandAddProMembership(t *testing.T) {
	received := CommandAddProMembership
	expected := "ADD_PRO_MEMBERSHIP"
	if expected != received {
		t.Errorf("Expected %s, Received %s", expected, received)
	}
}
func TestCommandApplyCoupon(t *testing.T) {
	received := CommandApplyCoupon
	expected := "APPLY_COUPON"
	if expected != received {
		t.Errorf("Expected %s, Received %s", expected, received)
	}
}
func TestCommandPrintBill(t *testing.T) {
	received := CommandPrintBill
	expected := "PRINT_BILL"
	if expected != received {
		t.Errorf("Expected %s, Received %s", expected, received)
	}
}
