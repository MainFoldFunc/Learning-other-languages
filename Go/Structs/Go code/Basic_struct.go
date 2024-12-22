package main

import "fmt"

// Define the car struct outside the main function so it's accessible globally
type car struct {
	Model      string
	Color      string
	Plate      int
	FrontWheel Wheel // Corrected field name and type
	BackWheel  Wheel // Corrected field name and type
}

type Wheel struct {
	Mark   string
	Radius float64
}

func params(p car) {
	// Using fmt.Printf instead of fmt.Print for proper formatting
	fmt.Printf("Car model is: %v\n", p.Model)
	fmt.Printf("Car plate is: %v\n", p.Plate)
	fmt.Printf("Car color is: %v\n", p.Color)
	fmt.Printf("Front wheel mark: %v, radius: %.2f\n", p.FrontWheel.Mark, p.FrontWheel.Radius)
	fmt.Printf("Back wheel mark: %v, radius: %.2f\n", p.BackWheel.Mark, p.BackWheel.Radius)
}

func main() {
	var exit string
	// Initialize the struct with values
	car_1 := car{
		Model: "Tesla A",
		Color: "Black",
		Plate: 43212,
		FrontWheel: Wheel{
			Mark:   "Chlemain",
			Radius: 21.3,
		},
		BackWheel: Wheel{
			Mark:   "Chlemain",
			Radius: 21.3,
		},
	}

	// Pass the car struct to the params function
	params(car_1)

	// Wait for user input before exiting
	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

