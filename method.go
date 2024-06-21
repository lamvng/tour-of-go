package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Method for Vertex type
// Receiver argument for method "v Vertex": Between the "func" keyword and the method name Abs()
// Methods are function for a type
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// "Normal" function
// Vertex instance as argument
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receiver
// The referenced struct instance is changed
// The method will modify the receiver
// Hence pointer receiver is more common than value receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Receiver without pointer
// The underlying instance is not changed
func (v Vertex) ScaleWithoutPointerReference(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	vertex := Vertex{3, 4}
	fmt.Println(vertex)
	fmt.Println(vertex.AbsMethod()) // Method attached to instance of the struct
	fmt.Println(AbsFunc(vertex))    // Struct instance passed as argument

	vertex.Scale((10))
	fmt.Println(vertex) // Values are changed

	vertex.ScaleWithoutPointerReference((10))
	fmt.Println(vertex) // Values are not changed. The func modifies a copy of the object
}
