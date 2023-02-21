package accounts

import "start_go/Bank/clients"

type SavingsAccount struct {
	Holder                     clients.Holder
	Agency, Account, Operation int
	balance                    float64
}

func (c *SavingsAccount) withdraw(value float64) string {
	getbalance := value > 0 && value <= c.balance
	if getbalance {
		c.balance -= value
		return "Saque efetuado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingsAccount) deposit(valueDeposit float64) (string, float64) {
	if valueDeposit > 0 {
		c.balance += valueDeposit
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Valor inválido", c.balance
	}
}

func (c *SavingsAccount) getBalance() float64 {
	return c.balance
}
