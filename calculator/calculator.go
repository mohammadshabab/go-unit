package calculator

import "errors"

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCalculator(minimumPurchaseAmount, discountAmount int) (*DiscountCalculator, error) {
	if minimumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not be zero")
	}
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}, nil
}
func (dc *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > dc.minimumPurchaseAmount {
		multiplier := purchaseAmount / dc.minimumPurchaseAmount
		return purchaseAmount - (dc.discountAmount * multiplier)
	}
	return purchaseAmount
}
