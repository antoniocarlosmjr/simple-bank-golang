package main

import (
	"fmt"

	"simple-bank-golang/entities"
	"simple-bank-golang/services"
)

func main() {
	userJoao := entities.User{"Joao", "1234", "Developer"}
	accountJoao := entities.Account{userJoao, 100, 2001, 0}

	var accountAntonio *services.Account
	accountAntonio = new(services.Account)
	accountAntonio.User.Name = "Antonio"
	accountAntonio.Deposit(100.50)

	var accountLuana *services.Account
	accountLuana = new(services.Account)
	accountLuana.User.Name = "Luana"
	accountAntonio.Deposit(50)

	fmt.Println(accountJoao)

	fmt.Println(accountAntonio.GetBalance())
	fmt.Println(accountAntonio.Withdraw(10))
	fmt.Println(accountAntonio.GetBalance())
	fmt.Println(accountAntonio.Deposit(100))
	fmt.Println(accountAntonio.GetBalance())

	fmt.Println(accountAntonio.Transfer(50, accountLuana))
	fmt.Println(*accountAntonio, *accountLuana)
}
