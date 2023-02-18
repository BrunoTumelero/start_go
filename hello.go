package main

import (
	"fmt"
	"os"
)

func main() {
	nome := startProgram()

	menu()

	comando := readComand()

	//if comando == 1 {
	//	fmt.Println("Monitorando...")
	//} else if comando == 2 {
	//	fmt.Println("Aqui estão todos os logs", nome)
	//} else if comando == 0 {
	//	fmt.Println("Estou me retirando ser supremo")
	//} else {
	//	fmt.Println("Não consigo executar essa ordem ", nome)
	//}

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Aqui estão todos os logs", nome)
	case 0:
		fmt.Println("Estou me retirando ser supremo")
		os.Exit((0))
	default:
		fmt.Println("Não consigo executar essa ordem ", nome)
		os.Exit((-1))
	}
}

func startProgram() string {
	var nome string = "Lord Ainz"
	fmt.Println("Olá, ", nome)
	fmt.Println("Qual a sua ordem?")

	return nome
}

func readComand() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Ver Logs")
	fmt.Println("0 - Sair")
}
