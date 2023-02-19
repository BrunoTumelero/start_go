package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {

		nome := startProgram()

		menu()

		comando := readComand()

		switch comando {
		case 1:
			start_monitorning()
		case 2:
			fmt.Println("Aqui estão todos os logs", nome)
		case 0:
			fmt.Println("Estou me retirando ser supremo")
			os.Exit((0))
		default:
			fmt.Println("Não consigo executar essa ordem ", nome)
		}
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

func start_monitorning() {
	fmt.Println("Monitorando...")
	site := "https://app.skeel.com.br/"
	response, err := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "=========================> Ativo")
	} else {
		fmt.Println("Site:", site, "está fora do ar", err)
	}
}
