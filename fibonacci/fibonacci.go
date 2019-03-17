package main

import (
	"fmt"
	"time"
)

func main() {
	go spin()
	fmt.Printf("\r%d\n", fibonacci(64))
}

func fibonacci(n uint) uint {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func spin() {
	const moon string = `🌕🌖🌗🌘🌑🌒🌓🌔`
	for {
		for _, c := range moon {
			time.Sleep(8 * 16 * time.Millisecond)
			fmt.Printf("\r%c", c)
		}
	}
}
