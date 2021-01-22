// WHAT I UNDERSTAND AS SINGLE RESPONSIBILITY PRINCIPLE

package main

import (
	"fmt"
	"math"
)

// Circle struct
// Has one property radius
// Has 2 method calcArea and Describe
type circle struct {
	radius float64
}

func (c circle) calcArea() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) describe() string {
	return fmt.Sprintf("Shape : Circle\nRadius : %.5f", c.radius)
}

// Suare struct
// Has one property length
// Has 2 method calcArea and Describe
type square struct {
	length float64
}

func (s square) calcArea() float64 {
	return s.length * s.length
}
func (s square) describe() string {
	return fmt.Sprintf("Shape : Sqare\nLength : %.5f", s.length)
}

// Sqare and cicle both have descibe and calcArea
// So bring them under a common interface
// Then use will be able to use other struct by making use of this interface
// This interface is like a superclass
// So any other struct who uses this interface does not need to know how the descibe and
// calcArea is implemented
// They will be able to use the describe and calcArea methond
// For example our outputter struct
type shapes interface {
	describe() string
	calcArea() float64
}

// This struct only has method
// Method to output any shapes in different format
type outputter struct {
}

// Method for outputting the shapes in text format
// This method makes use of the shape interface so that it always knows
// the shape has two method called descibe and calcArea
// This func makes use of those tow functions to output the shapes in text format
// Similarly outputter can have other methods such as JSONOut, XML OUT etc
func (out outputter) textOut(shape shapes) string {
	description := shape.describe()
	area := shape.calcArea()
	return fmt.Sprintf("%s\nArea : %.5f\n", description, area)

}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	output := outputter{}
	// Outputter receives variable of type shape
	// circle and sqare implements shapes interface
	// So circle and square are also type of shape
	// So where we can pass shapes we can pass circles or sqare also
	// Note that in GoLang interfaces are implemented implicitly
	fmt.Println(output.textOut(c))
	fmt.Println(output.textOut(s))
}
