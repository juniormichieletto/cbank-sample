package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardID string
	CreateAt     time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreateAt = time.Now()
	return t
}
