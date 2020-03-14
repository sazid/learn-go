package main

import (
	"fmt"
	"math"
)

// Shape : interface for any structs which has an area
// Define the interface type
type Shape interface {
	area() float64
}

// Circle : defines a circlular object
type Circle struct {
	x, y, radius float64
}

// Rectangle : defines a rectangular object
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	// Get the area of the different shapes using the Shape interface
	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
