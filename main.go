package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	menu()
}

func menu() {
	var chooser int

	fmt.Println("Initialize app - 1")
	fmt.Println("Logs - 2")
	fmt.Println("Exit - 0")

	fmt.Scan(&chooser)

	switch chooser {
	case 1:
	case 2:
	case 0:
		os.Exit(0)
	default:
		fmt.Errorf("command not valid")
		menu()
	}
}
