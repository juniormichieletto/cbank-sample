package usecase

import (
	"time"

	"github.io/junirmichieletto/codebank/domain"
	"github.io/junirmichieletto/codebank/dto"
)

type TransactionUseCase struct {
	TransactionRepository domain.TransactionRepository
}

func NewTransactionUseCase(repository domain.TransactionRepository) TransactionUseCase {
	return TransactionUseCase{TransactionRepository: repository}
}

func (u TransactionUseCase) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {

	creditCard := u.hydrateCreditCard(transactionDto)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.newTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)
	err = u.TransactionRepository.SaveTransaction(*t, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	transactionDto.ID = t.ID
	transactionDto.CreatedAt = t.CreatedAt
	if err != nil {
		return domain.Transaction{}, err
	}
	return *t, nil
}

func (u TransactionUseCase) hydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV
	return creditCard
}

func (u TransactionUseCase) newTransaction(transaction dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transaction.Amount
	t.Store = transaction.Store
	t.Description = transaction.Description
	t.CreatedAt = time.Now()
	return t
}
