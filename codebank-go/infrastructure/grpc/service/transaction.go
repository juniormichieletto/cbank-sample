package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.io/junirmichieletto/codebank/dto"
	"github.io/junirmichieletto/codebank/infrastructure/grpc/pb"
	"github.io/junirmichieletto/codebank/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.TransactionUseCase
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, request *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto := dto.Transaction{
		Name:            request.GetCreditCard().GetName(),
		Number:          request.GetCreditCard().GetNumber(),
		ExpirationMonth: request.GetCreditCard().GetExpirationMonth(),
		ExpirationYear:  request.CreditCard.GetExpirationYear(),
		Amount:          request.GetAmount(),
		CVV:             request.CreditCard.GetCvv(),
		Store:           request.GetStore(),
		Description:     request.GetDescription(),
	}
	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionDto)
	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	if !transaction.IsApproved() {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "Transaction rejected rejected by the bank")
	}
	return &empty.Empty{}, nil
}
