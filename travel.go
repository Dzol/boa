package main

import "fmt"
import "math"

func main() {
	var a, v, s float64
	var t float64

	fmt.Print("Please enter a value for acceleration: ")
	_, _ = fmt.Scan(&a)

	fmt.Print("Please enter a value for initial velocity: ")
	_, _ = fmt.Scan(&v)

	fmt.Print("Please enter a value for initial displacement: ")
	_, _ = fmt.Scan(&s)

	f := displacement(a, v, s)

	fmt.Print("Please enter a value for time: ")
	_, _ = fmt.Scan(&t)

	fmt.Printf("displacement after time %f = %f\n", t, f(t))
}

func displacement(a, v_0, s_0 float64) func(float64) float64 {
	return func(t float64) float64 {
		s := (a * math.Pow(t, 2)/ 2.0) + (v_0 * t) + s_0
		return s
	}
}
