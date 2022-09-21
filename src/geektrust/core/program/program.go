package program

type Program struct {
	Category ProgramCategory
	Quantity int
}

// Returns the total pro-membership discount to be applied on selected program(s).
// The discount is proportional to no of programs (Quantity field).
// The pro-membership is applicable only when the student has added pro-membership fee.
func (p Program) ProMembershipDiscount(hasProMemberShip bool) float64 {
	if hasProMemberShip {
		category := p.Category
		return category.Fee() * category.ProMembershipDiscount() * float64(p.Quantity)
	}
	return 0
}