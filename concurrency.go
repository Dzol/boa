package main

import "fmt"

func main() {
	c := make(chan uint)
	f := func(x uint) {
		for i := uint(0); i < x; i++ {
			c <- i
		}
		close(c)
	}
	go f(6)
	for i := range c {
		fmt.Println(i)
	}
}
