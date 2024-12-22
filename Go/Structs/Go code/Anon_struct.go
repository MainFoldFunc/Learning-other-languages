package main

import "fmt"

func from_struct(p struct{
	Model string
	Color string
	Year  int
}) {
	fmt.Printf("Model of car: %v\n", p.Model)
	fmt.Printf("Color of car: %v\n", p.Color)
	fmt.Printf("Year of production of car: %v\n", p.Year)
}

func main() {
	var exit string
	// Initialize the anonymous struct with values
	car_1 := struct {
		Model string
		Color string
		Year  int
	}{
		Model: "Tesla A",
		Color: "White",
		Year:  2020,
	}

	// Pass the anonymous struct to the function
	from_struct(car_1)

	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

