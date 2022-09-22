/*
Package program - package is central repository for managing the program categories and corresponding cost. Pro membership discounts also be defined here.
*/
package program

// ProgramCategory - data type to define programme category.
type ProgramCategory int64

// 3 types of program categories
const (
	CategoryUnknown ProgramCategory = iota
	CategoryCertification
	CategoryDegree
	CategoryDiploma
)

// Individual programmes pro membership discount (in percentage)
//
// DIPLOMA - 1% discount
//
// CERTIFICATION - 2% discount
//
// DEGREE - 3% discount
const (
	ProDiscountDiploma       = 0.01
	ProDiscountCertification = 0.02
	ProDiscountDegree        = 0.03
)

// ProgramCost - data type to define the programme cost
type ProgramCost float64

// cost of the program based on category
//
// CERTIFICATION - Rs. 3000
//
// DEGREE - Rs. 5000
//
// DIPLOMA - Rs 2500
const (
	CostCertification ProgramCost = 3000
	CostDegree        ProgramCost = 5000
	CostDiploma       ProgramCost = 2500
)

// Returns the human readable version for program category enum
func (c ProgramCategory) String() string {
	switch c {
	case CategoryCertification:
		return "CERTIFICATION"
	case CategoryDegree:
		return "DEGREE"
	case CategoryDiploma:
		return "DIPLOMA"
	default:
		return "unknown category"
	}
}

// Fee - Returns the individual program category amount
func (c ProgramCategory) Fee() float64 {
	switch c {
	case CategoryCertification:
		return float64(CostCertification)
	case CategoryDegree:
		return float64(CostDegree)
	case CategoryDiploma:
		return float64(CostDiploma)
	default:
		return 0
	}
}

// ProMembershipDiscount - Returns the discount (percentage) applicable on each and individual program category
func (c ProgramCategory) ProMembershipDiscount() float64 {
	switch c {
	case CategoryCertification:
		return ProDiscountCertification
	case CategoryDegree:
		return ProDiscountDegree
	case CategoryDiploma:
		return ProDiscountDiploma
	}
	return 0
}
