package main

import (
	"math/rand"
)

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   float64
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000), // Random ID for demonstration
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000000)), // Random account number
		Balance:   0.0,                          // Initial balance
	}
}
