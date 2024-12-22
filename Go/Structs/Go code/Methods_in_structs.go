package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

func params(r rectangle) {
	fmt.Printf("The width of this rectangle: %v\n", r.width)
	fmt.Printf("The height of this rectangle is: %v\n", r.height)
}

func area(r rectangle) float64 {
	return r.width * r.height
}

func main() {
	var exit string
	var width float64
	var height float64

	fmt.Print("What is the height of this rectangle: \n")
	fmt.Scan(&height)
	fmt.Print("What is the width of this rectangle: \n")
	fmt.Scan(&width)


	// Initialize the rectangle struct
	rect_1 := rectangle{
		width:  width,
		height: height,
	}

	// Call the params function to print width and height
	params(rect_1)

	// Calculate the area of the rectangle and store it in area_r
	area_r := area(rect_1)

	// Print the area of the rectangle
	fmt.Printf("The area of this rectangle is %v\n", area_r)

	// Prompt to exit
	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

