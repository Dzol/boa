package main

import "fmt"

func main() {
	fmt.Println(greet("Joe"))
}

func greet(s string) string {
	return "Hey" + " " + s + "!"
}
