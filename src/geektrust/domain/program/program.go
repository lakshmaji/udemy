package program

type Program struct {
	Category ProgramCategory
	Qty      int
}

// Returns the total pro-membership discount to be applied on selected program(s).
// The discount is proportional to no of programs (Qty field).
// The pro-membership is applicable only when the student has added pro-membership fee.
func (p Program) ProMembershipDiscount(hasProMemberShip bool) float64 {
	if hasProMemberShip {
		category := p.Category
		return category.Fee() * category.ProMembershipDiscount() * float64(p.Qty)
	}
	return 0
}
