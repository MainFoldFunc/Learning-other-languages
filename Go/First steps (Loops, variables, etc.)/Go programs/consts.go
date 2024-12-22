package main

import "fmt"

func main(){
	const name = "Jon"
	var exit string
	fmt.Println("You can't re asign constant: ", name)
	fmt.Println("Press e to exit...")
	fmt.Scan(&exit)
}	
	
