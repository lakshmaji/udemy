package coupon

type Coupons []Coupon

func (list Coupons) Max() Coupon {
	if len(list) == 0 {
		return ""
	}
	if len(list) == 1 {
		return list[0]
	}

	var max Coupon

	for _, coupon := range list {
		if coupon.Percentage() > max.Percentage() {
			max = coupon
		}
	}
	return max
}
