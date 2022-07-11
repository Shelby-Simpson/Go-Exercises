package Wallet

import (
	"BankApplication/BankAccount"
	"BankApplication/Entity"
	"fmt"
)

type Wallet struct {
	ID       string
	Accounts []BankAccount.BankAccount
	Owner    Entity.Entity
}

func (w *Wallet) DisplayAccounts() {
	for _, val := range w.Accounts {
		if val.AccountType == "checking" {
			fmt.Println(val)
		}
	}
	for _, val := range w.Accounts {
		if val.AccountType == "investment" {
			fmt.Println(val)
		}
	}
	for _, val := range w.Accounts {
		if val.AccountType == "savings" {
			fmt.Println(val)
		}
	}
}

func (w *Wallet) Balance() float64 {
	var balance float64
	for _, val := range w.Accounts {
		balance += val.Balance
	}
	return balance
}

func (w *Wallet) Wire(accountNum string, b *BankAccount.BankAccount, amount float64) {
	var account *BankAccount.BankAccount
	for i := range w.Accounts {
		if w.Accounts[i].AccountNo == accountNum {
			account = &(w.Accounts[i])
		}
	}
	if account.Balance >= amount {
		account.Balance -= amount
		(*b).Balance += amount
	}
}
