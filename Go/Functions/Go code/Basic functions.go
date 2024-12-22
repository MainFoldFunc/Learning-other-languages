package main

import (
	"fmt"
)

func help(name string) {
	fmt.Printf("Hello %v\n", name)
}

func main() {
	var name string
	var exit string

	fmt.Print("What is your name: \n")
	fmt.Scan(&name) // Corrected typo
	help(name)

	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

