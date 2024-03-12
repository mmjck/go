package main

import "fmt"

func learnTypes() {

	s := "Learn Go"
	s2 := `1500s, when an unknown printer took a galley
		like Aldus PageMaker including versions of Lorem Ipsum.
	`

	fmt.Println(s)
	fmt.Println(s2)

	g := 'Σ'
	f := 3.14159

	c := 3 + 4i

	fmt.Println(g)
	fmt.Println(f)
	fmt.Println(c)

	var u uint = 7

	var pi float32 = 22.7 / 7

	n := byte('\n')

	fmt.Println(u, pi, n)
	var a4 [4]int
	a3 := [...]int{1, 2, 3}

	s4 := make([]int, 4)

	var d2 [][]float64

	bs := []byte("a slice")

	fmt.Println(a4, a3)
	fmt.Println(bs, d2, s4)

}
