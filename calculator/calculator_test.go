package calculator

import (
	"testing"
)

/*
//depricated with code change keeping for reference
func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if amount != 130 {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(40)

	// if amount != 40 {
	// 	//t.Logf() only print the message but still need to call t.Fail() to make a test fail
	// 	//so need to call 2 methods for test failure
	// 	t.Logf("expected 50, got %v ", amount)
	// 	t.Fail()
	// }
	if amount != 40 {
		//Go testing package provides a combination of log and fail method called Error() or Errorf
		t.Errorf("expected 50, got %v. failed because the discount was not suppose to apply ", amount)
	}
}

func TestDiscountMultipliedByTwoApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(200)

	if amount != 160 {
		t.Errorf("expected 160, got %v ", amount)
	}
}

func TestDiscountMultipliedByThreeApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(350)

	if amount != 290 {
		t.Errorf("expected 290, got %v ", amount)
	}
}


//Table driven test {get rid of repeating code}
//We can use Table driven test when the structure is the same
//if the test code is different make another test otherwise will need to write lot of ifs
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchageAmount int
		discount              int
		purchaseAmaount       int
		expectedAmount        int
	}
	testCases := []testCase{
		{minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 150, expectedAmount: 130},
		{minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 200, expectedAmount: 160},
		{minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 350, expectedAmount: 290},
		{minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 40, expectedAmount: 50},
	}
	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchageAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmaount)
		if amount != tc.expectedAmount {
			t.Errorf("expected %v, got %v ", tc.expectedAmount, amount)
		}
	}
}
*/
//the problem above is hard to know what test case failed because to the testing package
//we have only a single test to run
//There is a feature called sub tests
//With sub tests we can inform to the testing package that each row in our for loop
//is a different test case.
//Another good part of subtest is we can run a specific test case
//go test -v -run="TestDiscountCalculatorSubTest/Should_apply_20" ./...

func TestDiscountCalculatorSubTest(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchageAmount int
		discount              int
		purchaseAmaount       int
		expectedAmount        int
	}
	testCases := []testCase{
		{name: "Should apply 20", minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 150, expectedAmount: 130},
		{name: "Should apply 40", minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 200, expectedAmount: 160},
		{name: "Should apply 60", minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 350, expectedAmount: 290},
		{name: "Should not apply", minimumPurchageAmount: 100, discount: 20, purchaseAmaount: 50, expectedAmount: 50},
		// {name: "zero minimum purchase amount", minimumPurchageAmount: 0, discount: 20, purchaseAmaount: 50, expectedAmount: 50},
	}
	for _, tc := range testCases {
		// Run func receives a name and second parameter is a function
		//And we can further improve this by giving a name
		t.Run(tc.name, func(t *testing.T) {
			// t.Run("Discount test", func(t *testing.T) {
			calculator, err := NewDiscountCalculator(tc.minimumPurchageAmount, tc.discount)
			if err != nil {
				//here test execution will not stop because t.Errorf is a combination of Fail+log
				//so everything after this will fail
				// //t.Errorf("could not instantiate the calculator %s ", err.Error())
				//t.FailNow() this will stop the execution
				// //t.FailNow()
				//but t.FailNow() wil not print the message
				//t.Fatalf() is a combination of FailNow()+log
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
		_, err := NewDiscountCalculator(tc.minimumPurchageAmount, tc.discount)
		if err == nil {
			t.Fatalf("should not create the calculator with zero purchase amount")
		}
	}
}
