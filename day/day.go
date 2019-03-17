package main

import "fmt"

type Day int

const (
	M Day = iota
	Tu
	W
	Th
	F
	Sa
	Su
)

var x = [...]string{
	M:  "Monday",
	Tu: "Tuesday",
	W:  "Wednesday",
	Th: "Thursday",
	F:  "Friday",
	Sa: "Saturday",
	Su: "Sunday",
}

func (d Day) String() string {
	return x[d]
}

func main() {
	fmt.Println(Sa)
	fmt.Println(Su)
}
