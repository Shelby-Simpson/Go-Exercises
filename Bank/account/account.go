package account

import (
	"fmt"
	"time"
)

type Account struct {
	Number               int
	Balance              float64
	Name                 string
	TransactionDateArray []time.Time
}

func (a *Account) Deposit(amt float64) {
	a.Balance += amt
	a.TransactionDateArray = append(a.TransactionDateArray, time.Now())
	fmt.Println(a.TransactionDateArray)
}

func (a *Account) Withdraw(amt float64) {
	if amt > a.Balance {
		fmt.Printf("You cannot withdraw %v.\n", amt)
	} else {
		a.Balance -= amt
		a.TransactionDateArray = append(a.TransactionDateArray, time.Now())
		fmt.Println(a.TransactionDateArray)
	}
}
