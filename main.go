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
	fmt.Println(accountAntonio.withdraw(10))
	fmt.Println(accountAntonio.balance)
	fmt.Println(accountAntonio.deposit(100))
	fmt.Println(accountAntonio.balance)
}

func (account *Account) withdraw(value float64) string {
	hasBalance := value <= account.balance
	if hasBalance && value > 0 {
		account.balance -= value
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente ou valor informado é inválido"
}

func (account *Account) deposit(value float64) string {
	if value > 0 {
		account.balance += value
		return "Depósito realizado com sucesso"
	}

	return "Valor do depósito é informado é inválido"
}
