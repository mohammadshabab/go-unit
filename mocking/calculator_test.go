package mocking

import (
	"testing"
)

//go test go-unit\calculator     running particular package
//go test -v ./... to run all  v for verbose
type DiscountRepositoryMock struct{}

func (dr DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscountCalculatorSubTest(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchageAmount int
		purchaseAmaount       int
		expectedAmount        int
	}
	testCases := []testCase{
		{name: "Should apply 20", minimumPurchageAmount: 100, purchaseAmaount: 150, expectedAmount: 130},
		{name: "Should apply 40", minimumPurchageAmount: 100, purchaseAmaount: 200, expectedAmount: 160},
		{name: "Should apply 60", minimumPurchageAmount: 100, purchaseAmaount: 350, expectedAmount: 290},
		{name: "Should not apply", minimumPurchageAmount: 100, purchaseAmaount: 50, expectedAmount: 50},
		// {name: "zero minimum purchase amount", minimumPurchageAmount: 0, discount: 20, purchaseAmaount: 50, expectedAmount: 50},
	}
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			//database.DiscountRepository() //we will have to mock this repo as it needs a db connection
			discountRepositoryMock := DiscountRepositoryMock{}
			calculator, err := NewDiscountCalculator(tc.minimumPurchageAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("could not instantiate the calculator %s ", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmaount)
			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v ", tc.expectedAmount, amount)
			}
		})

	}
}

//to get rid of if err
func TestDiscountCalculatorShouldFailWithZeroMinimumAmount(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchageAmount int
		discount              int
		purchaseAmaount       int
		expectedAmount        int
	}
	testCases := []testCase{
		{name: "zero minimum purchase amount", minimumPurchageAmount: 0, discount: 20, purchaseAmaount: 50, expectedAmount: 50},
	}
	for _, tc := range testCases {
		discountRepositoryMock := DiscountRepositoryMock{}
		_, err := NewDiscountCalculator(tc.minimumPurchageAmount, discountRepositoryMock)
		if err == nil {
			t.Fatalf("should not create the calculator with zero purchase amount")
		}
	}
}
