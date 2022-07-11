package BankAccount

import (
	"BankApplication/Entity"
	"fmt"
)

type BankAccount struct {
	AccountNo    string
	AccountOwner Entity.Entity
	Balance      float64
	InterestRate float64
	AccountType  string
}

func (a *BankAccount) Withdraw(amount float64) {
	if a.Balance < amount {
		fmt.Println("You cannot withdraw this much.")
	} else {
		a.Balance -= amount
	}
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) ApplyInterest() {
	if a.AccountOwner.EntityType == "Individual" {
		switch a.AccountType {
		case "checking":
			a.Balance += a.Balance * .01
		case "investment":
			a.Balance += a.Balance * .02
		case "savings":
			a.Balance += a.Balance * .05
		}
	} else if a.AccountOwner.EntityType == "Business" {
		switch a.AccountType {
		case "checking":
			a.Balance += a.Balance * .005
		case "investment":
			a.Balance += a.Balance * .01
		case "savings":
			a.Balance += a.Balance * .02
		}
	}
}

func (a *BankAccount) Wire(b *BankAccount, amount float64) {
	if a.Balance >= amount {
		a.Balance -= amount
		(*b).Balance += amount
	}
}
