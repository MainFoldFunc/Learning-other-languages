package main

import "fmt"

func main(){
	var exit string
	var arr [10]int
	
	slice := arr[0 : 9]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	slice_of_primes := primes[0: 3]

	new_slice := make([]int, 5)

	fmt.Printf("This is array of zeros: \n%v\n", arr)
	fmt.Printf("This is slice made from arr: \n%v\n", slice)
	fmt.Printf("This is array of primes: \n%v\n", primes)
	fmt.Printf("This is the slice of primes: \n%v\n", slice_of_primes)
	fmt.Printf("This is new slice without list: \n%v\n", new_slice)
	fmt.Printf("THe length of this slcie is \n%v\n", len(new_slice))
	fmt.Printf("The max capcity of this slice is: \n%v\n", cap(new_slice))

	fmt.Print("\nPress e to exit...")
	fmt.Scan(&exit)
}
