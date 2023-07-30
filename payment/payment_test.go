package payment

import (
	"testing"

	"github.com/mohammadshabab/go-unit/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AttemptHistory struct {
	mock.Mock
}

func (a *AttemptHistory) IncrementFailure(user entity.User) error {
	args := a.Called(user)

	return args.Error(0)
}
func (a *AttemptHistory) CountFailures(user entity.User) (int, error) {
	args := a.Called(user)
	return args.Int(0), args.Error(1)
}

type GatewayMock struct {
	mock.Mock
}

func (g *GatewayMock) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {
	args := g.Called(user, creditCard)

	return args.Bool(0), args.Error(1)
}
func (g *GatewayMock) Pay(crediCard entity.CreditCard, amount int) error {
	args := g.Called(crediCard, amount)
	return args.Error(0)
}

func TestShouldHaveSuccessfullAuthorization(t *testing.T) {
	user := entity.User{}
	creditCard := entity.CreditCard{}
	attemptHistory := &AttemptHistory{}
	attemptHistory.On("CountFailures", user).Return(1, nil)
	gateway := &GatewayMock{}
	gateway.On("IsAuthorized", user, creditCard).Return(true, nil)
	paymentService := NewPaymentService(attemptHistory, gateway)

	isAuthorized, err := paymentService.IsAuthorized(user, creditCard)
	if err != nil {
		t.Fatal(err.Error())
	}
	attemptHistory.AssertNotCalled(t, "IncrementFailure", user)
	assert.True(t, isAuthorized)
}

func TestShouldHaveAFailureAuthorization(t *testing.T) {
	user := entity.User{}
	creditCard := entity.CreditCard{}
	attemptHistory := &AttemptHistory{}
	attemptHistory.On("CountFailures", user).Return(1, nil)
	attemptHistory.On("IncrementFailure", user).Return(nil)
	gateway := &GatewayMock{}
	gateway.On("IsAuthorized", user, creditCard).Return(false, nil)
	paymentService := NewPaymentService(attemptHistory, gateway)

	isAuthorized, err := paymentService.IsAuthorized(user, creditCard)
	if err != nil {
		t.Fatal(err.Error())
	}
	attemptHistory.AssertCalled(t, "IncrementFailure", user)
	assert.False(t, isAuthorized)
}

func TestShouldHaveAForcedFailureAuthorization(t *testing.T) {
	//dummy A dummy is when you pass some object just to satisfy the argument list of
	//some method call
	user := entity.User{}
	creditCard := entity.CreditCard{}

	attemptHistory := &AttemptHistory{} //like a stub because we are just returning value
	attemptHistory.On("CountFailures", user).Return(6, nil)

	gateway := &GatewayMock{} // like a mock because we are asserting behavour

	paymentService := NewPaymentService(attemptHistory, gateway)

	isAuthorized, err := paymentService.IsAuthorized(user, creditCard)
	if err != nil {
		t.Fatal(err.Error())
	}
	gateway.AssertNotCalled(t, "IsAuthorized", user, creditCard)
	assert.False(t, isAuthorized)
}
