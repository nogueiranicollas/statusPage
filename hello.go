package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	introducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func introducao() {
	var nome = "Nicollas"
	var versao float32 = 1.2

	fmt.Println("Olá, Sr. ", nome)
	fmt.Println("Este programa esta na versão : ", versao)

}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O Site: ", site, " foi carregado com sucesso")
	} else {
		fmt.Println("O Site: ", site, " Está com problema ", resp.StatusCode)
	}
	//fmt.Println(resp)

}
