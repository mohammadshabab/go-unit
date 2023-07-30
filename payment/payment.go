package payment

import (
	"github.com/mohammadshabab/go-unit/entity"

	"github.com/mohammadshabab/go-unit/database"
)

type PaymentService struct {
	attemptHistoryRepository database.AttemptHistory
	gateway                  Gateway
}

func NewPaymentService(attemptHistoryRepository database.AttemptHistory, gateway Gateway) *PaymentService {
	return &PaymentService{
		attemptHistoryRepository: attemptHistoryRepository,
		gateway:                  gateway,
	}
}

func (p *PaymentService) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {
	totalAttempts, err := p.attemptHistoryRepository.CountFailures(user)
	if err != nil {
		return false, err
	}
	if totalAttempts > 5 {
		return false, nil
	}

	isAuthorized, err := p.gateway.IsAuthorized(user, creditCard)
	if err != nil {
		return false, err
	}

	if !isAuthorized {
		err := p.attemptHistoryRepository.IncrementFailure(user)
		if err != nil {
			return false, err
		}
	}
	return isAuthorized, nil
}
