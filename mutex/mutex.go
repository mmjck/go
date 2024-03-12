package main

import (
	"fmt"
	"sync"
)

var total = 0

func count(wg *sync.WaitGroup) {
	total++
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count(&wg)
	}

	wg.Wait()

	fmt.Println("total: ", total)
}
