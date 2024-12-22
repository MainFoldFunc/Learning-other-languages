package main

import "fmt"

func calc(num_1 float64, num_2 float64, sign string) float64 {
	switch sign {
	case "+":
		return num_1 + num_2
	case "-":
		return num_1 - num_2
	case "*":
		return num_1 * num_2
	case "/":
		if num_2 == 0 {
			fmt.Println("Error: Division by zero")
			return 0
		}
		return num_1 / num_2
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func main() {
	var num_1 float64
	var num_2 float64
	var operator string
	var res float64
	var exit string

	fmt.Print("Enter number 1: ")
	fmt.Scan(&num_1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&num_2)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	res = calc(num_1, num_2, operator)

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operator. Please use +, -, *, or /")
	} else {
		fmt.Printf("%v %v %v = %v\n", num_1, operator, num_2, res)
	}

	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

