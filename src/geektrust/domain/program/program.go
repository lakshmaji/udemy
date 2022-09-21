package program

type Program struct {
	Category ProgramCategory
	Qty      int
}

func (p Program) ProMembershipDiscount(hasProMemberShip bool) float64 {
	if hasProMemberShip {
		category := p.Category
		return category.Fee() * category.ProMembershipDiscount() * float64(p.Qty)
	}
	return 0
}
