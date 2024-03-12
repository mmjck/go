package main

import "fmt"

func for2() {
	j := 0
	for j < 5 {
		fmt.Println("Value of j:", j)
		j++
	}

	k := 0
	for {
		fmt.Println("Value of k:", k)
		k++

		if k == 5 {
			break
		}
	}
}
