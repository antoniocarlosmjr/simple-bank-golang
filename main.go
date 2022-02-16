package main

import "fmt"

type Account struct {
	owner   string
	agency  int
	account int
	balance float64
}

/* Aplicação para gerenciar contas correntes */
func main() {
	fmt.Println(Account{
		owner:   "",
		agency:  0,
		account: 0,
		balance: 0,
	})
}
