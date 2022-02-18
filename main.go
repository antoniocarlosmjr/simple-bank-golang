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

	fmt.Println(*accountAntonio)
}
