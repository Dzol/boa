package main

import "fmt"

func main() {
	i := 0
	fmt.Printf("i = %d\n", i)                   // i = 0
	defer fmt.Printf("defer i + 2 = %d\n", i+2) // i + 2 = 2
	fmt.Println("defer @ i + 2 = 2")
	i += 1
	fmt.Printf("i = %d\n", i) // i = 1
}
