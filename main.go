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
	contaAntonio := Account{
		owner:   "Antonio Martins",
		agency:  10,
		account: 123456,
		balance: 100.50,
	}
	fmt.Println(contaAntonio)
}
