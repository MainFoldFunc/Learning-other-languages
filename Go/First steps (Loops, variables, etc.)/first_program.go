package main

import "fmt"

func main(){
	name := ""
	output := ""
	fmt.Println("Hello world")
	fmt.Println("Press what is your name: ")
	fmt.Scan(&name)
	fmt.Println("Your name is:", name)
	fmt.Println("Press e to exit")
	fmt.Scan(&output)
	
}
