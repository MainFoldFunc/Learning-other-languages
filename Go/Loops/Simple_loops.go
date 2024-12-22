package main

import "fmt"

func add_x(num, howmuch_to_add int) int{
	for i := 0; i < howmuch_to_add; i++{
		num++
	}
	return num
}

func main(){
	var exit string
	var num int
	var how_much_add int

	fmt.Print("What is the num: ")
	fmt.Scan(&num)

	fmt.Print("How much would you like to add: ")
	fmt.Scan(&how_much_add)

	num = add_x(num, how_much_add)

	fmt.Printf("The number after adding %v is %v", how_much_add, num)

	fmt.Print("\nPress e to exit...")
	fmt.Scan(&exit)
}
