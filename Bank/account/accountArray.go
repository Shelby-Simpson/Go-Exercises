package account

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type AccountArray struct {
	Array        []*Account
	mostRecentID int
}

func (accountArray *AccountArray) CreateAccount(name string, initialBalance float64) {
	accountArray.Array = append(accountArray.Array, &Account{accountArray.mostRecentID + 1, initialBalance, name, []time.Time{}})
	accountArray.mostRecentID += 1
}

func (accountArray *AccountArray) FindAccount(name string) Account {
	for i := range accountArray.Array {
		if (*accountArray.Array[i]).Name == name {
			return (*accountArray.Array[i])
		}
	}
	return Account{}
}

func (accountArray *AccountArray) SortByName() {
	sort.Slice(accountArray.Array, func(i, j int) bool {
		return strings.ToLower((*accountArray.Array[i]).Name) < strings.ToLower((*accountArray.Array[j]).Name)
	})
}

func (accountArray *AccountArray) SortByBalance() {
	sort.Slice(accountArray.Array, func(i, j int) bool {
		return (*accountArray.Array[i]).Balance < (*accountArray.Array[j]).Balance
	})
}

func (accountArray *AccountArray) Print() {
	for _, val := range accountArray.Array {
		fmt.Println("Account number:", val.Number)
		fmt.Println("Account name:", val.Name)
		fmt.Println("Account balance:", val.Balance)
		fmt.Println("Account transaction history:", val.TransactionDateArray)
		fmt.Println()
	}
}
