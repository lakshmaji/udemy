package program

type ProgramCategory int64

const (
	CategoryUnknown ProgramCategory = iota
	CategoryCertification
	CategoryDegree
	CategoryDiploma
)

const (
	ProDiscountDiploma       = 0.01
	ProDiscountCertification = 0.02
	ProDiscountDegree        = 0.03
)

type ProgramCost float64

const (
	CostCertification ProgramCost = 3000
	CostDegree        ProgramCost = 5000
	CostDiploma       ProgramCost = 2500
)

func (c ProgramCategory) String() string {
	switch c {
	case CategoryCertification:
		return "CERTIFICATION"
	case CategoryDegree:
		return "DEGREE"
	case CategoryDiploma:
		return "DIPLOMA"
	}
	return "unknown category"
}

func (c ProgramCategory) Fee() float64 {
	switch c {
	case CategoryCertification:
		return float64(CostCertification)
	case CategoryDegree:
		return float64(CostDegree)
	case CategoryDiploma:
		return float64(CostDiploma)
	}
	return 0
}

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
