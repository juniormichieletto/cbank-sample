package usecase

import (
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

	return domain.Transaction{}, nil
}
