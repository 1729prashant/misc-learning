package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line


func updateBalance(c *customer, t transaction) error {
	
	if t.transactionType == transactionDeposit {
		c.balance += t.amount
		return nil
	}

	if t.transactionType == transactionWithdrawal {
		if c.balance - t.amount < 0 {
			return errors.New("insufficient funds")
		} else {
			c.balance -= t.amount
			return nil
		}
	}

	return errors.New("unknown transaction type")

}
