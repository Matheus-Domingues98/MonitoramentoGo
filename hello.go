package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			// Indica que a saida do programada ocorreu bem
			os.Exit(0)
		default:
			fmt.Println("Não reconheço esse comando.")
			// Indica que ocorreu algum erro no codigo ou coia inesperada
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {

	nome := "Matheus"
	versao := 1.1
	fmt.Println("Ola Sr.", nome)
	fmt.Println("Esse programa esta na versao", versao)
	fmt.Println("----------")
}

func lerComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando... ")

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
			fmt.Println(" ")
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("--------------------")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro ao testar site:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problema. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return sites
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

	arquivo, err := os.OpenFile("log.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")
	
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de log:", err)
		return
	}
	fmt.Println(string(arquivo))

}
