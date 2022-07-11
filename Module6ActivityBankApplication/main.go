// Shelby Simpson
// Module 6 Activity: Bank Application
package main

import (
	"BankApplication/BankAccount"
	"BankApplication/Entity"
	"BankApplication/Wallet"
)

func main() {
	wallet1 := Wallet.Wallet{
		ID:       "1",
		Accounts: make([]BankAccount.BankAccount, 0),
		Owner: Entity.Entity{
			ID:         1,
			Address:    "123 Street St",
			EntityType: "Individual",
		},
	}
	wallet1.Accounts = append(wallet1.Accounts, BankAccount.BankAccount{
		AccountNo:    "1",
		AccountOwner: wallet1.Owner,
		Balance:      50,
		InterestRate: .01,
		AccountType:  "checking",
	})
	wallet2 := Wallet.Wallet{
		ID:       "2",
		Accounts: make([]BankAccount.BankAccount, 0),
		Owner: Entity.Entity{
			ID:         2,
			Address:    "123 Road Rd",
			EntityType: "Business",
		},
	}
	wallet2.Accounts = append(wallet2.Accounts, BankAccount.BankAccount{
		AccountNo:    "2",
		AccountOwner: wallet2.Owner,
		Balance:      100,
		InterestRate: .005,
		AccountType:  "checking",
	})
	wallet1.Wire(wallet1.Accounts[0].AccountNo, &(wallet2.Accounts[0]), 25)
	wallet1.DisplayAccounts()
	wallet2.DisplayAccounts()
	for i := range wallet1.Accounts {
		wallet1.Accounts[i].ApplyInterest()
	}
	wallet2.Accounts[0].Deposit(50)
	wallet1.DisplayAccounts()
	wallet2.DisplayAccounts()
}
