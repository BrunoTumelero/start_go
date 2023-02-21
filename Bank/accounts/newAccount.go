package accounts

import "start_go/Bank/clients"

type CurrentAccount struct {
	Holder  clients.Holder
	Agency  int
	Account int
	Balance float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	getbalance := value > 0 && value <= c.Balance
	if getbalance {
		c.Balance -= value
		return "Saque efetuado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CurrentAccount) Deposit(valueDeposit float64) (string, float64) {
	if valueDeposit > 0 {
		c.Balance += valueDeposit
		return "Depósito realizado com sucesso", c.Balance
	} else {
		return "Valor inválido", c.Balance
	}
}

func (c *CurrentAccount) Transfer(valueTransfer float64, destination *CurrentAccount) bool {
	if valueTransfer < c.Balance && valueTransfer > 0 {
		c.Balance -= valueTransfer
		destination.Deposit(valueTransfer)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) getBalance() float64 {
	return c.Balance
}
