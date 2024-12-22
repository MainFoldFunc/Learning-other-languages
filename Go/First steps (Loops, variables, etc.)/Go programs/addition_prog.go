package main

import "fmt"

func main(){
	var num_1 int
	var num_2 int
	var score int
	var exit string

	fmt.Println("Enter a first number: ")
	fmt.Scan(&num_1)
	fmt.Println("Enter a second number: ")
	fmt.Scan(&num_2)

	score = num_1 + num_2
	fmt.Println("number 1 + number 2 = ", score)

	fmt.Println("Press e to exit")
	fmt.Scan(&exit)

}
