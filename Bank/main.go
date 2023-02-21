package main

import (
	"fmt"
	"start_go/Bank/accounts"
	"start_go/Bank/clients"
)

func main() {
	clientMarcos := clients.Holder{"Marcos", "123.456.789.50", "Vendedor"}
	fmt.Println(clientMarcos)
	marcosAccount := accounts.CurrentAccount{}
	marcosSavings := accounts.SavingsAccount{Holder: clientMarcos}
	fmt.Println(marcosAccount)
	fmt.Println(marcosSavings)

}
