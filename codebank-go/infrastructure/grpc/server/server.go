package server

import (
	"log"
	"net"

	"github.io/junirmichieletto/codebank/infrastructure/grpc/pb"
	"github.io/junirmichieletto/codebank/infrastructure/grpc/service"
	"github.io/junirmichieletto/codebank/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	TransactionUseCase usecase.TransactionUseCase
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (s GRPCServer) Start() error {
	address := "0.0.0.0:50052"
	list, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Could not listem tcp port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = s.TransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	return grpcServer.Serve(list)
}
