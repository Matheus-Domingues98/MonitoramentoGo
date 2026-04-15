package main

import (
	"fmt"
	"net/http"
	"os"
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
	sites := []string{"https://www.alura.com.br",
		"https://www.caelum.com.br",
		"https://www.uol.com.br"}

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

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problema. Status Code:", resp.StatusCode)
	}
}
