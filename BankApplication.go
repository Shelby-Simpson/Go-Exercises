// Shelby Simpson
// Bank Application
package main

import "fmt"

type BankAccount struct {
	accountNo    string
	accountOwner entity
	balance      float64
	interestRate float64
	accountType  string
}

type entity struct {
	ID         int
	Address    string
	entityType string
}

type Wallet struct {
	ID       string
	accounts []BankAccount
	owner    entity
}

func (a *BankAccount) withdraw(amount float64) {
	if a.balance < amount {
		fmt.Println("You cannot withdraw this much.")
	} else {
		a.balance -= amount
	}
}

func (a *BankAccount) deposit(amount float64) {
	a.balance += amount
}

func (a *BankAccount) applyInterest() {
	if a.accountOwner.entityType == "Individual" {
		switch a.accountType {
		case "checking":
			a.balance += a.balance * .01
		case "investment":
			a.balance += a.balance * .02
		case "savings":
			a.balance += a.balance * .05
		}
	} else if a.accountOwner.entityType == "Business" {
		switch a.accountType {
		case "checking":
			a.balance += a.balance * .005
		case "investment":
			a.balance += a.balance * .01
		case "savings":
			a.balance += a.balance * .02
		}
	}
}

func (a *BankAccount) wire(b *BankAccount, amount float64) {
	if a.balance >= amount {
		a.balance -= amount
		(*b).balance += amount
	}
}

func (e *entity) changeAddress(newAddr string) {
	e.Address = newAddr
}

func (w *Wallet) displayAccounts() {
	for _, val := range w.accounts {
		if val.accountType == "checking" {
			fmt.Println(val)
		}
	}
	for _, val := range w.accounts {
		if val.accountType == "investment" {
			fmt.Println(val)
		}
	}
	for _, val := range w.accounts {
		if val.accountType == "savings" {
			fmt.Println(val)
		}
	}
}

func (w *Wallet) balance() float64 {
	var balance float64
	for _, val := range w.accounts {
		balance += val.balance
	}
	return balance
}

func (w *Wallet) wire(accountNum string, b *BankAccount, amount float64) {
	var account *BankAccount
	for i := range w.accounts {
		if w.accounts[i].accountNo == accountNum {
			account = &(w.accounts[i])
		}
	}
	if account.balance >= amount {
		account.balance -= amount
		(*b).balance += amount
	}
}

func main() {
	wallet1 := Wallet{
		ID:       "1",
		accounts: make([]BankAccount, 0),
		owner: entity{
			ID:         1,
			Address:    "123 Street St",
			entityType: "Individual",
		},
	}
	wallet1.accounts = append(wallet1.accounts, BankAccount{
		accountNo:    "1",
		accountOwner: wallet1.owner,
		balance:      50,
		interestRate: .01,
		accountType:  "checking",
	})
	wallet2 := Wallet{
		ID:       "2",
		accounts: make([]BankAccount, 0),
		owner: entity{
			ID:         2,
			Address:    "123 Road Rd",
			entityType: "Business",
		},
	}
	wallet2.accounts = append(wallet2.accounts, BankAccount{
		accountNo:    "2",
		accountOwner: wallet2.owner,
		balance:      100,
		interestRate: .005,
		accountType:  "checking",
	})
	wallet1.wire(wallet1.accounts[0].accountNo, &(wallet2.accounts[0]), 25)
	wallet1.displayAccounts()
	wallet2.displayAccounts()
	for i := range wallet1.accounts {
		wallet1.accounts[i].applyInterest()
	}
	wallet2.accounts[0].deposit(50)
	wallet1.displayAccounts()
	wallet2.displayAccounts()
}
