package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}
type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func getArea(shape Shape) float64 {
	return shape.Area()
}
