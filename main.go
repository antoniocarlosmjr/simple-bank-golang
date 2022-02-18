package main

import (
	"fmt"
)

type Account struct {
	owner   string
	agency  int
	account int
	balance float64
}

/* Aplicação para gerenciar contas correntes */
func main() {
	var accountAntonio *Account
	accountAntonio = new(Account)
	accountAntonio.owner = "Antonio"
	accountAntonio.balance = 100.50

	var accountLuana *Account
	accountLuana = new(Account)
	accountLuana.owner = "Luana"
	accountLuana.balance = 50.

	fmt.Println(accountAntonio.balance)
	*accountAntonio = withdrawAccount(*accountAntonio, 100)
	fmt.Println(accountAntonio.balance)
}

func withdrawAccount(account Account, value float64) Account {
	if value <= account.balance {
		account.balance = account.balance - value
	} else {
		fmt.Println("Impossível sacar valor, pois não possui saldo suficiente")
	}

	return account
}
