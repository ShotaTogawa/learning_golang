package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(sh shape) {
	x := sh.area()
	fmt.Println(x)
}


func main() {
	c := circle{radius: 3.0}
	s := square{length: 3}

	info(c)
	info(s)
}
