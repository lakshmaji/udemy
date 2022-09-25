package coupon

// Coupon - a model for defining coupon code.
type Coupon string

// Available coupons
const (
	CouponDealG20 Coupon = "DEAL_G20"
	CouponDealG5  Coupon = "DEAL_G5"
	CouponB4G1    Coupon = "B4G1"
)

const (
	// CouponDealG20MarginAmount - DEAL_G20 coupon purchase programme value
	CouponDealG20MarginAmount = 10000
)

const (
	// CouponDealG5MarginCount - DEAL_G5 coupon minimum program count value
	CouponDealG5MarginCount int = 2
	// CouponB4G1MarginCount - B4G1 coupon  minimum program count value
	CouponB4G1MarginCount int = 4
)

const (
	// DealG20 - This coupon can be applied if the purchased programmes value is Rs.10,000/- or above. It provides a 20% discount on the total programme cost. The coupon needs to be applied explicitly to get a discount.
	// This coupon percentage value
	dealG20 float64 = 0.2 // 20 / 100
	// DealG5 - This coupon can only be applied if there are a minimum of 2 programmes being purchased. It provides a 5% discount on the total programme cost. The coupon needs to be applied explicitly to get a discount.
	// This coupon percentage value
	dealG5 float64 = 0.05 // 5 / 100
)

// Percentage - returns the percentage discount (number) configured over given coupon (if applicable).
func (c Coupon) Percentage() float64 {
	switch c {
	case CouponDealG20:
		return dealG20
	case CouponDealG5:
		return dealG5
	default:
		return 0
	}
}
