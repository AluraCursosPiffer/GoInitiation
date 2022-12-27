package main

import (
	"fmt"
	metodos "github.com/AluraCursosPiffer/GoInitiation/src"
	"os"
)

func main() {
	metodos.ExibeIntroducao()

	for {
		metodos.ExibeMenu()

		comando := metodos.LeComando()

		switch comando {
		case 1:
			metodos.IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			metodos.ImprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}
