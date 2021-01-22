// Liscov Substitution principle
// This principle is strongly related to inheritance
// In Golang we do not have inheritance
// So It is difficult to recreate the principle easily

// Liscov Substitution Principle says that if a method can be
// called on some type (class)
// Then that method should also work on it's subtypes

// modified LSP for golang
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.

package main

import "fmt"

// Create a super type called quadrangle (চতুর্ভুজ)
type quadrangle interface {
	getWidth() int
	setWidth(width int)
	getHeight() int
	setHeight(height int)
}

// Now create a rectangle and implement the interface
// Note : interfaces are implemented implicitly in go
type rectangle struct {
	width, height int
}

// Implement getWitdht
func (r *rectangle) getWidth() int {
	return r.width
}

// Implement setWidth
func (r *rectangle) setWidth(width int) {
	r.width = width
}

// Implement getHeight
func (r *rectangle) getHeight() int {
	return r.height
}

// Implement getHeight
func (r *rectangle) setHeight(height int) {
	r.height = height
}

// Now I create another struct square which also implements quadrangle

type square struct {
	rectangle
}

func newSqare(side int) *square {
	sq := square{}
	sq.width = side
	sq.height = side
	return &sq
}

func (sq *square) getHeight() int {
	return sq.height
}
func (sq *square) setHeight(height int) {
	sq.height = height
	sq.width = height
}

func (sq *square) getWidth() int {
	return sq.width
}

func (sq *square) setWidth(width int) {
	sq.width = width
	sq.height = width
}

func useQuadrangle(quad quadrangle) {

	// Modify the height of the quadrangle
	width := quad.getWidth()

	quad.setHeight(10)

	expectedArea := 10 * width
	actualArea := quad.getWidth() * quad.getHeight()
	// Expectation : expected area and actual area should be the same
	fmt.Println("Expected Area : ", expectedArea, "Actual Area : ", actualArea)

}

func main() {

	rc := &rectangle{2, 3}
	useQuadrangle(rc) // Area should be 20
	sq := newSqare(5)
	useQuadrangle(sq) // Area should be 50

	// For the sqare the useQuadrangle does not work as expected
	// Hence the useQuadrangle function is breaking the LSP
	// For the sqare the useQuadrangle does not work as expected
	// Hence the useQuadrangle function is breaking the LSP
	// Becasue when we created the useQuadrangle function we made
	// assumption that amything which is a type of quadrangle
	// setting its height will only set its height
	// But in this case when the sq sub type is passed the assumption
	// is broken

}
