package main

import (
	"fmt"
	"math"
)

const π = 3.14

func main() {
	fmt.Println(maximum(π, 4.0, 8.0, 15.0, 16.0, 23.0, 42.0))
}

func maximum(i ...float64) float64 {
	maximum := math.Inf(-1)
	for _, v := range i {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}
