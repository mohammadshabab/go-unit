package mocking

import (
	"errors"

	"github.com/mohammadshabab/go-unit/database"
)

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountRepository    database.Repository
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountRepository database.Repository) (*DiscountCalculator, error) {
	if minimumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not be zero")
	}
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}
func (dc *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > dc.minimumPurchaseAmount {
		multiplier := purchaseAmount / dc.minimumPurchaseAmount

		discount := dc.discountRepository.FindCurrentDiscount()

		return purchaseAmount - (discount * multiplier)
	}
	return purchaseAmount
}
