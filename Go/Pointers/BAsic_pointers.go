package main

import "fmt"

func main(){
	var exit string
	var x int

	x = 10
	y := &x

	fmt.Printf("The pointer to the x is: %v", y)

	*y = 11
	fmt.Printf("\nThe pointer after change is %v", *y)
	fmt.Printf("\nThe original value is: %v", x)

	fmt.Print("\nPress e to exit... ")
	fmt.Scan(&exit)
}
