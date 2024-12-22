package main

import "fmt"

func main(){
	var num_1 float64
	var num_2 float64
	var score float64
	var exit string

	fmt.Print("What is first number: \n")
	fmt.Scan(&num_1)
	fmt.Print("What is second number: \n")
	fmt.Scan(&num_2)

	score = num_1 + num_2

	fmt.Printf("%.2f + %.2f = %.2f\n", num_1, num_2, score)
	
	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)


}
