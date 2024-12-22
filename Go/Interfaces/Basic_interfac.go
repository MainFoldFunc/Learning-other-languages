package main

import "fmt"

// Define the shape interface
type shape interface {
	Area() float64
}

// Define rect type with width and height
type rect struct {
	width  float64
	height float64
}

// Define circle type with radius
type circle struct {
	radius float64
}

// Implement the Area method for rect
func (r rect) Area() float64 {
	return r.width * r.height
}

// Implement the Area method for circle
func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Function to print the area of a shape
func print_area(s shape, sh string) {
	fmt.Printf("The area of shape %v is around: %.4v\n", sh, s.Area())
}

func main() {
	var exit string
	var height float64
	var width float64
	var radius float64

	// Get rectangle dimensions
	fmt.Print("What is the height of the rectangle: \n")
	fmt.Scan(&height)
	fmt.Print("What is the width of the rectangle: \n")
	fmt.Scan(&width)

	// Create a rectangle instance
	r := rect{
		width:  width,
		height: height,
	}

	// Print the area of the rectangle
	print_area(r, "rectangle")

	// Get circle radius
	fmt.Print("What is the radius of the circle: \n")
	fmt.Scan(&radius)

	// Create a circle instance
	c := circle{
		radius: radius,
	}

	// Print the area of the circle
	print_area(c, "circle")

	// Wait for user input before exiting
	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

