package main

import "fmt"

func main(){
	var bool_1 float64
	var bool_2 float64
	var score float64
	var exit string

	fmt.Println("Enter first bool number:")
	fmt.Scan(&bool_1)
	fmt.Println("Enter second bool number:")
	fmt.Scan(&bool_2)

	score = bool_1 + bool_2
	fmt.Println("The addition of those two numbers is: ", score)

	fmt.Println("Press e to exit")
	fmt.Scan(&exit)

}
