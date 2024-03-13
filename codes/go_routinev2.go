package main

import (
	"fmt"
	"time"
)

func sayTo(s string, done chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	done <- "Finished"
}

func mainn() {
	done := make(chan string)

	go sayTo("world", done)
	fmt.Println(<-done)

}
