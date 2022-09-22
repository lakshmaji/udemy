package coupon

// Coupons - This specifically deals with the list of coupons applied by the user.
// This is basically a field from Cart instance.
//
// If 2 or more coupons are applied, the higher value coupon needs to be considered.
type Coupons []Coupon

// Pick - Returns one coupon among list of coupons applied (if any),
// otherwise return empty string which represents no coupon
func (list Coupons) Pick() Coupon {
	couponCount := len(list)
	// When there are no coupons applied
	if couponCount == 0 {
		return list.None()
	}
	// When there is one coupon applied
	if couponCount == 1 {
		return list.First()
	}
	return list.Max()
}

// First - Select the first coupon from the list
func (list Coupons) First() Coupon {
	return list[0]
}

// None - No coupon is applied
func (list Coupons) None() Coupon {
	return ""
}

// Max - When there are more than one coupon applied, pick the coupon with maximum discount
func (list Coupons) Max() Coupon {
	var max Coupon
	for _, coupon := range list {
		if coupon.Percentage() > max.Percentage() {
			max = coupon
		}
	}
	return max
}
