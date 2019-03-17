package main

import "fmt"
import "encoding/json"
import "math"

type Point struct {
	X int
	Y int
	z int
}

type GeometryPointer interface {
	distance2O() float64
}

type CartesianPt struct {
	X int
	Y int
}

func (p CartesianPt) distance2O() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type ManhattanPt struct {
	X int
	Y int
}

func (p ManhattanPt) distance2O() float64 {
	return float64(p.X + p.Y)
}

type ThreeDPt struct {
	X int
	Y int
	Z int
}

func (p ThreeDPt) distance2O() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

func main() {
	var q Point
	var x, y, z GeometryPointer

	p := Point{3, 2, 1}

	x = CartesianPt{3, 2}
	y = ManhattanPt{3, 2}
	z = ThreeDPt{3, 2, 1}

	o, _ := json.Marshal(p)
	e := json.Unmarshal(o, &q)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(q)

	f(x)
	f(y)
	f(z)
}

func f(p GeometryPointer) {
	fmt.Printf("distance from origin: %f\n", p.distance2O())
}
