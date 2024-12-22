package main

import "fmt"

// Define the shape interface
type shape interface {
	Area() float64
}

// Define the structs for various shapes
type rect struct {
	width  float64
	height float64
}
type square struct {
	side float64
}
type triangle struct {
	base   float64
	height float64
}
type circle struct {
	radius float64
}
type trapezoid struct {
	base_d float64
	base_u float64
	height float64
}

// Implement the Area method for each shape
func (r rect) Area() float64 {
	return r.width * r.height
}
func (s square) Area() float64 {
	return s.side * s.side
}
func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}
func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (tr trapezoid) Area() float64 {
	return ((tr.base_d + tr.base_u) * tr.height) / 2
}

// Function to print the area of a shape
func print_area(r shape, sh string) {
	fmt.Printf("The area of %v is around: %.4f\n", sh, r.Area())
}

func main() {
	var exit string
	var height, width, side, base, radius, base_u, base_d float64

	// Rectangle
	fmt.Print("Enter the height of the rectangle: ")
	fmt.Scan(&height)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&width)
	rect := rect{width: width, height: height}
	print_area(rect, "rectangle")

	// Square
	fmt.Print("Enter the side of the square: ")
	fmt.Scan(&side)
	square := square{side: side}
	print_area(square, "square")

	// Triangle
	fmt.Print("Enter the base of the triangle: ")
	fmt.Scan(&base)
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scan(&height)
	triangle := triangle{base: base, height: height}
	print_area(triangle, "triangle")

	// Circle
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)
	circle := circle{radius: radius}
	print_area(circle, "circle")

	// Trapezoid
	fmt.Print("Enter the upper base of the trapezoid: ")
	fmt.Scan(&base_u)
	fmt.Print("Enter the lower base of the trapezoid: ")
	fmt.Scan(&base_d)
	fmt.Print("Enter the height of the trapezoid: ")
	fmt.Scan(&height)
	trapezoid := trapezoid{base_u: base_u, base_d: base_d, height: height}
	print_area(trapezoid, "trapezoid")

	// Exit prompt
	fmt.Print("Press e to exit...")
	fmt.Scan(&exit)
}

