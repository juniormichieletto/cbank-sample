package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreateAt        time.Time
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreateAt = time.Now()
	return c
}
