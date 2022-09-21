package coupon

type Coupons []Coupon

// Returns one coupon among list of coupons applied (if any),
// otherwise return empty string which represents no coupon
func (list Coupons) Pick() Coupon {
	// When there are no coupons applied
	if len(list) == 0 {
		return ""
	}
	// When there is one coupon applied
	if len(list) == 1 {
		return list[0]
	}

	// When there are more than one coupon applied, pick the coupon with maximum discount
	var max Coupon
	for _, coupon := range list {
		if coupon.Percentage() > max.Percentage() {
			max = coupon
		}
	}
	return max
}
