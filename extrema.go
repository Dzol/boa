package main

import "fmt"

const π = 3.14

func main() {
	fmt.Println(maximum(4, 8, 15, 16, 23, 42))
}

func maximum(i ...uint) uint {
	maximum := uint(0)
	for _, v := range i {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}
