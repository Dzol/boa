package main

import "fmt"

type Grade int

const (
	A Grade = iota
	B
	C
	D
	E
	F
)

func main() {
	x := A
	fmt.Printf("x :: %T = %v\n", x, x)
}
