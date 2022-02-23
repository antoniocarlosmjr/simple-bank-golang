package main

import (
	"fmt"

	"simple-bank-golang/services"
)

func main() {
	var accountAntonio *services.Account
	accountAntonio = new(services.Account)
	accountAntonio.User.Name = "Antonio"
	accountAntonio.Balance = 100.50

	var accountLuana *services.Account
	accountLuana = new(services.Account)
	accountLuana.User.Name = "Luana"
	accountLuana.Balance = 50.

	fmt.Println(accountAntonio.Balance)
	fmt.Println(accountAntonio.Withdraw(10))
	fmt.Println(accountAntonio.Balance)
	fmt.Println(accountAntonio.Deposit(100))
	fmt.Println(accountAntonio.Balance)

	fmt.Println(accountAntonio.Transfer(50, accountLuana))
	fmt.Println(*accountAntonio, *accountLuana)
}
