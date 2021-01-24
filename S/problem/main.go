// Here the describe method does 2 things
// calculates the area
// Prints the area
// Hence the single responsibility principle is violated

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) describe() {
	fmt.Printf("circle area: %f\n", math.Pi*c.radius*c.radius)
}

type square struct {
	length float64
}

func (s square) describe() {
	fmt.Printf("square area: %f\n", s.length*s.length)
}

func main() {
	c := circle{radius: 5}
	c.describe()

	s := square{length: 7}
	s.describe()
}
