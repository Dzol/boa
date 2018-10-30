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
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (d Day) String() string {
	return x[d]
}

func main() {
	fmt.Println(Sa)
	fmt.Println(Su)
}
