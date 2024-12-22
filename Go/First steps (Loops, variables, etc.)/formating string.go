package main

import "fmt"

func main(){
	var age int
	var name string
	var exit string

	fmt.Print("What is your name: \n")
	fmt.Scan(&name)
	fmt.Print("How old are you: \n")
	fmt.Scan(&age)
	
	fmt.Printf("Hey %s you are %d years old\n", name, age)

	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}
