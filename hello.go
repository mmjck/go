package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	beyondHello()
	learnTypes()

	p, q := learnMemory()

	fmt.Println(*p, *q)

	m := map[string]int{"três": 3, "quatro": 4}
	m2 := map[string]bool{"f": false, "t": true}

	fmt.Println(m2)
	fmt.Println(m)

	m["um"] = 1

	fmt.Println(m)

	_, _ = learnMemory()

	learnFlowControl()

	rectangle := Rectangle{width: 5, height: 3}
	circle := Circle{radius: 4}

	// Calculate and print areas
	fmt.Println("Area of rectangle:", getArea(rectangle))
	fmt.Println("Area of circle:", getArea(circle))

}
