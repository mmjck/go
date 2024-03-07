package main

import (
	"fmt"
)

func learnFlowControl() {
	if true {
		fmt.Println("eu avisei-te")
	}

	if false {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

	x := 1

	switch x {
	case 0:
	case 1:
	case 2:
	}

	for i := 0; i < 10; i++ {
		fmt.Println("it", i)
	}

	for {
		break
		continue
	}

	if y := 1 + 3; y < 1 {
		x = y
	}

	xBig := func() bool {
		return x > 100
	}

	fmt.Println("xBig:", xBig())
	x /= 1e5
	fmt.Println("xBig:", xBig())
}
