package main

import "fmt"

func main() {
	x := []int{1,2,3}
	f(x)
	fmt.Println(x)
}

func f(s []int) {
	for i, v := range s {
		s[i] = 2 * v
	}
}
