package main

import (
	"fmt"
	"github.com/MainFoldFunc/Learning-other-languages/Go/Projects/First-Project/custpack"
)

func main() {
	// Declare a string variable to hold user input
	var str string

	// Ask the user for input
	fmt.Println("What is the string: ")
	fmt.Scan(&str)

	// Call the Reverse function from the custpack package and print the result
	fmt.Printf("The string reversed is: %s\n", custpack.Reverse(str))
}

