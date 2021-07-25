package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.io/junirmichieletto/codebank/infrastructure/grpc/server"
	"github.io/junirmichieletto/codebank/infrastructure/kafka"
	"github.io/junirmichieletto/codebank/infrastructure/repository"
	"github.io/junirmichieletto/codebank/usecase"
)

func main() {
	db := setupDb()
	defer db.Close()

	producer := setupKakfaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.TransactionUseCase {
	transactionRepository := repository.NewTransactionRepositoryDB(db)
	useCase := usecase.NewTransactionUseCase(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}

func serveGrpc(processTransactionUseCase usecase.TransactionUseCase) {
	grpcServer := server.NewGRPCServer()
	grpcServer.TransactionUseCase = processTransactionUseCase
	fmt.Println("serveGrpc is running")
	grpcServer.Start()
}

func setupKakfaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetProducer("host.docker.internal:9094")
	return producer
}

/*
func createCreditCardAndPersistForDBValidationPurposes(db *sql.DB) {
	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Junior"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDB(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}
*/
