package main

import "fmt"

type CurrentAccount struct {
	holder  string
	agency  int
	account int
	balance float64
}

func (c *CurrentAccount) withdraw(value float64) string {
	getBalance := value > 0 && value <= c.balance
	if getBalance {
		c.balance -= value
		return "Saque efetuado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	brunoAccount := CurrentAccount{
		holder:  "Bruno",
		agency:  123,
		account: 456321,
		balance: 200,
	}
	marcosAccount := CurrentAccount{"marcos", 787, 968574, 147.58}
	fmt.Println(brunoAccount)
	fmt.Println(marcosAccount)

	var accountjoao *CurrentAccount
	accountjoao = new(CurrentAccount)
	accountjoao.holder = "João"
	accountjoao.agency = 741
	accountjoao.account = 753159
	accountjoao.balance = 270.47

	fmt.Println(*accountjoao)
	fmt.Println(accountjoao.withdraw(2000))
	fmt.Println(*accountjoao)
}
