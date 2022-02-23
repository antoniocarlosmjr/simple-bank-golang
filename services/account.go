package services

import "simple-bank-golang/entities"

type Account struct {
	User    entities.User
	Agency  int
	Account int
	balance float64
}

func (account *Account) Withdraw(value float64) string {
	hasBalance := value <= account.balance
	if hasBalance && value > 0 {
		account.balance -= value
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente ou valor informado para depósito é inválido"
}

func (account *Account) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.balance += value
		return "Depósito realizado com sucesso", account.balance
	}

	return "Valor do depósito informado é inválido", account.balance
}

func (accountOrigin *Account) Transfer(value float64, accountTarget *Account) string {

	balanceAccountOrigin := accountOrigin.balance

	if value > 0 && value <= balanceAccountOrigin {
		accountOrigin.balance -= value
		accountTarget.Deposit(value)

		return "Transferência realizada com sucesso"
	}

	return "Valor é menor do que zero e conta de origem precisa de saldo em conta"
}

func (account *Account) GetBalance() float64 {
	return account.balance
}
