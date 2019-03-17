package main

import "fmt"

func main() {
	var f float64
	fmt.Print("Please enter a float: ")
	fmt.Scan(&f)
	fmt.Println(int(f))
}
