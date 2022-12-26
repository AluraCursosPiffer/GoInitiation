package main

import (
	"fmt"
	"os"
)

const monitoramentos = 2
const delay = 5

func main() {
	metodos.exibeIntroducao()

	for {
		metodos.exibeMenu()

		comando := metodos.leComando()

		switch comando {
		case 1:
			metodos.iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			metodos.imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}
