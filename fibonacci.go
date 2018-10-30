package main

import (
	"fmt"
	"time"
)

func main() {
	go spin()
	fmt.Printf("\r%d\n", fibonacci(36))
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
	for {
		for _, c := range `-\|/` {
			time.Sleep(64 * time.Millisecond); fmt.Printf("\r%c", c)
		}
	}
}
