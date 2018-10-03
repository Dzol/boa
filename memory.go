package main

import "fmt"

func main() {
	var y *int
	y = f()
	fmt.Printf("y: %d\n", y)
}

func f() *int {
	x := 1024
	return &x
}
