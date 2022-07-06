// Shelby Simpson
// Practice: Create A Struct
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type account struct {
	number               int
	balance              float64
	name                 string
	transactionDateArray []time.Time
}

func (a *account) deposit(amt float64) {
	a.balance += amt
	a.transactionDateArray = append(a.transactionDateArray, time.Now())
	fmt.Println(a.transactionDateArray)
}

func (a *account) withdraw(amt float64) {
	if amt > a.balance {
		fmt.Printf("You cannot withdraw %v.\n", amt)
	} else {
		a.balance -= amt
		a.transactionDateArray = append(a.transactionDateArray, time.Now())
		fmt.Println(a.transactionDateArray)
	}
}

func sortByName(accountArray []*account) {
	sort.Slice(accountArray, func(i, j int) bool {
		return strings.ToLower((*accountArray[i]).name) < strings.ToLower((*accountArray[j]).name)
	})
}

func soryByBalance(accountArray []*account) {
	sort.Slice(accountArray, func(i, j int) bool {
		return (*accountArray[i]).balance < (*accountArray[j]).balance
	})
}

func findAccount(accountArray []*account, name string) account {
	for i := range accountArray {
		if (*accountArray[i]).name == name {
			return (*accountArray[i])
		}
	}
	return account{}
}

func createAccount(accountArray []*account, name string, balance float64) {

}

func main() {
	acc1, acc2, acc3 := account{}, account{}, account{}
	accountArray := []*account{&acc1, &acc2, &acc3}

	for i := 0; i < len(accountArray); i++ {
		(*accountArray[i]).number = i + 1
		fmt.Println("Please enter the name for account", i+1)
		scanner := bufio.NewScanner(os.Stdin)
		if ok := scanner.Scan(); ok {
			(*accountArray[i]).name = scanner.Text()
		}
		fmt.Println("Please enter the initial balance for account", i+1)
		fmt.Scan(&accountArray[i].balance)
		fmt.Println("Account number: ", (*accountArray[i]).number)
		fmt.Println("Account name: ", (*accountArray[i]).name)
		fmt.Println("Account balance: ", (*accountArray[i]).balance)
	}

	sortByName(accountArray)

	for i := range accountArray {
		fmt.Println((*accountArray[i]).name)
	}

	soryByBalance(accountArray)

	for i := range accountArray {
		fmt.Println((*accountArray[i]).balance)
	}

	(*accountArray[0]).deposit(40)
	(*accountArray[0]).withdraw(30)
}
