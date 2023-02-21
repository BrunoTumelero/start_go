package accounts

import "start_go/Bank/clients"

type CurrentAccount struct {
	Holder  clients.Holder
	Agency  int
	Account int
	balance float64
}

func (c *CurrentAccount) withdraw(value float64) string {
	getbalance := value > 0 && value <= c.balance
	if getbalance {
		c.balance -= value
		return "Saque efetuado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CurrentAccount) deposit(valueDeposit float64) (string, float64) {
	if valueDeposit > 0 {
		c.balance += valueDeposit
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Valor inválido", c.balance
	}
}

func (c *CurrentAccount) transfer(valueTransfer float64, destination *CurrentAccount) bool {
	if valueTransfer < c.balance && valueTransfer > 0 {
		c.balance -= valueTransfer
		destination.deposit(valueTransfer)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) getBalance() float64 {
	return c.balance
}
