package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (this *Rectangle) area() float64 {
	return this.width * this.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{3, 4}
	fmt.Println(r1.area())

	c1 := Circle{5}
	fmt.Println(c1.area())
}
