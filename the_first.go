package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 1

func main() {
	for {

		nome := startProgram()

		menu()

		comando := readComand()

		switch comando {
		case 1:
			start_monitorning()
		case 2:
			printLogs()
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
	sites := readFile()

	for i := 0; i <= monitoring; i++ {
		for i := range sites {
			response, err := http.Get(sites[i])
			if response.StatusCode == 200 {
				fmt.Println("Site:", sites[i], "=========================> Ativo")
				saveLog(sites[i], true)
			} else {
				fmt.Println("Site:", sites[i], "=========================> Está com problemas", response.StatusCode, err)
				saveLog(sites[i], false)
			}
		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")
}

func readFile() []string {
	var sites []string

	arq, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	read := bufio.NewReader(arq)
	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	arq.Close()

	return sites
}

func saveLog(site string, status bool) {
	arq, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arq.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arq.Close()
}

func printLogs() {
	arq, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arq))
}
