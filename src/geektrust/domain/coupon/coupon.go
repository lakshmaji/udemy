package coupon

type Coupon string

const (
	CouponDealG20 Coupon = "DEAL_G20"
	CouponDealG5         = "DEAL_G5"
	CouponB4G1           = "B4G1"
)

// Amount margin
const (
	CouponDealG20MarginAmount = 10000
)

// Items count (no of programs) margin
const (
	CouponDealG5MarginCount = 2
	CouponB4G1MarginCount   = 4
)

const (
	DEAL_G20 = 0.2  // 20 / 100
	DEAL_G5  = 0.05 // 5 / 100
)

func (c Coupon) Percentage() float64 {
	switch c {
	case CouponDealG20:
		return DEAL_G20
	case CouponDealG5:
		return DEAL_G5
	}
	return 0
}
