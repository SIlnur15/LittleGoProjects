package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// Методы для Circle (реализует Shape)
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Методы для Rectangle (реализует Shape)
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var shape Shape = Circle{radius: 5}
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n",
		shape.Area(), shape.Perimeter()) // 78.54, 31.42

	shape = Rectangle{width: 3, height: 4}
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n",
		shape.Area(), shape.Perimeter()) // 12.00, 14.00
}
