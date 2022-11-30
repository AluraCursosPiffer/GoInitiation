package main

import "fmt"

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
}
