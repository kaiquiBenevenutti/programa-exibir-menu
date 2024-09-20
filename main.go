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

const monitoramentos = 3
const delay = 3

func main() {
	bemVindo()
	fmt.Println("")

	for {
		exibirMenu()
		comando := lerComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimirLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) //para informar ao computador que saiu com sucesso, sem erros
		default:
			fmt.Println("Tente novamente.")
			os.Exit(-1) //para informar ao computador que algo deu errado
		}
	}
}

func bemVindo() {
	nome := "Kaiqui"
	versao := 2.2
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // o & e para mostrar o caminho para onde a resposta deve ser salva, nesse caso o comando
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			resp, err := http.Get(site)
			if err != nil {
				fmt.Println("Ocorreu um erro:", err)
				return
			}
			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "foi carregado com sucesso!")
				registraLog(site, true)
			} else {
				fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
				registraLog(site, false)
			}
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func lerSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("registro.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {
	arquivo, err := ioutil.ReadFile("registro.txt")

	fmt.Println("Registro de Logs:")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
