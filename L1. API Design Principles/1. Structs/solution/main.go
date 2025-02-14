package main

import (
	"fmt"
)

// Create the Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Define Area method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define Perimeter method
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// Create a new Rectangle
	rect := Rectangle{
		Width:  5.0,
		Height: 10.0,
	}

	// Print area and perimeter
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}
