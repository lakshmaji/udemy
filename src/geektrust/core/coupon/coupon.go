package coupon

type Coupon string

// Available coupons
const (
	CouponDealG20 Coupon = "DEAL_G20"
	CouponDealG5         = "DEAL_G5"
	CouponB4G1           = "B4G1"
)

const (
	// DEAL_G20 coupon purchase programme value
	CouponDealG20MarginAmount = 10000
)

const (
	// DEAL_G5 coupon minimum program count value
	CouponDealG5MarginCount = 2
	// B4G1 coupon  minimum program count value
	CouponB4G1MarginCount = 4
)

const (
	// This coupon can be applied if the purchased programmes value is Rs.10,000/- or above. It provides a 20% discount on the total programme cost. The coupon needs to be applied explicitly to get a discount.
	DEAL_G20 = 0.2 // 20 / 100
	// This coupon can only be applied if there are a minimum of 2 programmes being purchased. It provides a 5% discount on the total programme cost. The coupon needs to be applied explicitly to get a discount.
	DEAL_G5 = 0.05 // 5 / 100
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
