package main

type Vehicle interface {
	move() string
}

type Car struct{}

func (c Car) move() string {
	return "Driving"
}

type Bicycle struct{}

func (b Bicycle) move() string {
	return "Cycling"
}
