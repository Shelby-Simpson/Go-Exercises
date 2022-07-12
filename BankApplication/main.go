// Shelby Simpson
// Bank Application
package main

func main() {
	wallet1 := Wallet{
		ID:       "1",
		Accounts: make([]BankAccount, 0),
		Owner: Entity{
			ID:         1,
			Address:    "123 Street St",
			EntityType: "Individual",
		},
	}
	wallet1.Accounts = append(wallet1.Accounts, BankAccount{
		AccountNo:    "1",
		AccountOwner: wallet1.Owner,
		Balance:      50,
		InterestRate: .01,
		AccountType:  "checking",
	})
	wallet2 := Wallet{
		ID:       "2",
		Accounts: make([]BankAccount, 0),
		Owner: Entity{
			ID:         2,
			Address:    "123 Road Rd",
			entityType: "Business",
		},
	}
	wallet2.Accounts = append(wallet2.Accounts, BankAccount{
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
