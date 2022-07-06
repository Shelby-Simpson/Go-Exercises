package main

import (
	"bank/account"
)

func main() {
	accountArray := account.AccountArray{}
	accountArray.CreateAccount("Shelby Simpson", 40.00)
	accountArray.CreateAccount("James Bond", 300.32)
	accountArray.Print()
	accountArray.SortByName()
	accountArray.Print()
}
