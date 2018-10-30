package main

import "fmt"

func main() {
	var y *int
	y = f()
	fmt.Printf("y @ 0x%0x\n", y)
	fmt.Printf("y value: %d\n", *y)
}

func f() *int {
	x := 1024
	return &x
}
