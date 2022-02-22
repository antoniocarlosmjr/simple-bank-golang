package services

type Account struct {
	Owner   string
	Agency  int
	Account int
	Balance float64
}

func (account *Account) Withdraw(value float64) string {
	hasBalance := value <= account.Balance
	if hasBalance && value > 0 {
		account.Balance -= value
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente ou valor informado para depósito é inválido"
}

func (account *Account) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.Balance += value
		return "Depósito realizado com sucesso", account.Balance
	}

	return "Valor do depósito informado é inválido", account.Balance
}

func (accountOrigin *Account) Transfer(value float64, accountTarget *Account) string {

	balanceAccountOrigin := accountOrigin.Balance

	if value > 0 && value <= balanceAccountOrigin {
		accountOrigin.Balance -= value
		accountTarget.Deposit(value)

		return "Transferência realizada com sucesso"
	}

	return "Valor é menor do que zero e conta de origem precisa de saldo em conta"
}
