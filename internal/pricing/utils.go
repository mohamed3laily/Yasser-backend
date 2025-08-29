package pricing

import "math"

func CalculateFinalPrice(basePrice int, discountPercent float64) int {
	if discountPercent <= 0 || basePrice <= 0 {
		return basePrice
	}
	
	discountMultiplier := 1.0 - (discountPercent / 100.0)
	discountedPrice := float64(basePrice) * discountMultiplier
	
	finalPrice := int(math.Round(discountedPrice))
	if finalPrice < 0 {
		return 0
	}
	
	return finalPrice
}

func CalculateDiscountAmount(basePrice int, discountPercent float64) int {
	if discountPercent <= 0 || basePrice <= 0 {
		return 0
	}
	
	discountAmount := float64(basePrice) * (discountPercent / 100.0)
	return int(math.Round(discountAmount))
}

func IsDiscounted(discountPercent float64) bool {
	return discountPercent > 0
}