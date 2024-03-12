package main

import (
	"fmt"
	"sync"
)

var total = 0

func count(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	total++
	m.Unlock()
	wg.Done()
}

func main() {
	var (
		wg sync.WaitGroup
		m  sync.Mutex
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count(&wg, &m)
	}

	wg.Wait()

	fmt.Println("total: ", total)
}
