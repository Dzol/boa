package main

import "fmt"

func main() {
	x := 3
	g := f(x)
	y := g(2)
	z := g(4)
	fmt.Println(y)
	fmt.Println(z)
}

func f(n int) func(int) int {
	return func(i int) int {
		n += i
		return n
	}
}
