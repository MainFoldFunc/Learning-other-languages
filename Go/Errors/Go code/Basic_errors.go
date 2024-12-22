package main

import (
	"errors"
	"fmt"
)

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	var num_1 float64
	var num_2 float64
	var res float64
	var err error
	var exit string

	fmt.Print("Enter number 1: ")
	fmt.Scan(&num_1)
	fmt.Print("\nEnter number 2: ")
	fmt.Scan(&num_2)

	// Call the division function with error handling
	res, err = division(num_1, num_2)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
	} else {
		fmt.Printf("\nNumber 1 divided by number 2 is: %v\n", res)
	}

	fmt.Print("\nPress e to exit...")
	fmt.Scan(&exit)
}

