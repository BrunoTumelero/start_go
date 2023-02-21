package main

import (
	"fmt"
	"start_go/Bank/accounts"
	"start_go/Bank/clients"
)

func PayBills(account verificationAccount, valueBill float64) {
	account.Withdraw(valueBill)
}

type verificationAccount interface {
	Withdraw(value float64) string
}

func main() {
	clientMarcos := clients.Holder{"Marcos", "123.456.789.50", "Vendedor"}
	fmt.Println(clientMarcos)
	marcosAccount := accounts.CurrentAccount{}
	marcosSavings := accounts.SavingsAccount{Holder: clientMarcos}
	fmt.Println(marcosAccount)
	fmt.Println(marcosSavings)
	marcosAccount.Deposit(1000)
	fmt.Println(marcosAccount.Balance)
	PayBills(&marcosAccount, 300)
	fmt.Println(marcosAccount.Balance)

}
