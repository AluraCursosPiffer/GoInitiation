package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Bem vindo! Escolha uma opção do menu abaixo:")
	menu()
}

func menu() {
	var chooser int

	for {
		fmt.Println("Initialize app - 1")
		fmt.Println("Logs - 2")
		fmt.Println("Exit - 0")

		fmt.Scan(&chooser)

		switch chooser {
		case 1:
			monitoramento()
		case 2:
		case 0:
			os.Exit(0)
		default:
			fmt.Errorf("command not valid")
		}
	}

}

func monitoramento() error {
	var site string

	fmt.Println("Digite o site que deseja monitorar:")
	fmt.Scan(&site)

	response, err := http.Get(site)

	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		fmt.Println("O site está funcionando perfeitamente")
	} else {
		fmt.Println(string(response.StatusCode))
	}

	return nil
}
