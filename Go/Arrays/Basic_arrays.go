package main

import "fmt"

func main(){
	var exit string
	var arr [10]int
	
	slice := arr[0 : 9]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	slice_of_primes := primes[0: 3]

	fmt.Printf("This is array of zeros: \n%v\n", arr)
	fmt.Printf("This is slice made from arr: \n%v\n", slice)
	fmt.Printf("This is array of primes: \n%v", primes)
	fmt.Printf("This is the slice of primes: \n%v\n", slice_of_primes)

	fmt.Print("\nPress e to exit...")
	fmt.Scan(&exit)
}
