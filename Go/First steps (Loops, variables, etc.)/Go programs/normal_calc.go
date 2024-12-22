package main

import "fmt"

func main() {
	var num_1 float64
	var num_2 float64
	var sign string
	var score float64
	var exit string

	fmt.Print("Enter number 1: \n")
	fmt.Scan(&num_1)
	fmt.Print("Enter number 2: \n")
	fmt.Scan(&num_2)
	fmt.Print("Enter operator: \n")
	fmt.Scan(&sign)

	if sign == "+" {
		score = num_1 + num_2
	} else if sign == "-" { // Corrected else placement
		score = num_1 - num_2
	} else if sign == "*" { // Corrected else placement
		score = num_1 * num_2
	} else if sign == "/" { // Corrected else placement
		score = num_1 / num_2
	} else {
		fmt.Print("No such operator\n")
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num_1, sign, num_2, score) // Fixed order of arguments in Printf

	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

