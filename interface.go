package main

import (
	"fmt"
	"math"
)

type I interface {
	Abs() float64
	PrintValues()
}

func main() {
	var i I

	// Nil interface value
	var emptyVertex *Vertex
	i = emptyVertex
	fmt.Println(i) // Null pointer
	i.PrintValues()

	// Defined interface value
	definedVertex := Vertex{3, 4}
	i = &definedVertex
	fmt.Println(definedVertex.Abs())
	i.PrintValues()

	// Type assertion
	vertexValue, isVertex := i.(*Vertex)
	fmt.Printf("%v %v\n", vertexValue, isVertex)
	// If isVertex == false, and isVertex is not defined in the call, the program will panic
	// Aka. Type assertion failed

	// Type switch
	// Multiple type assertion
	switch value := i.(type) {
	case *Vertex:
		fmt.Println(value, "Is Vertex.")
	default:
		fmt.Printf("%T\n", vertexValue)
	}

}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) PrintValues() {
	if v == nil {
		fmt.Printf("<nil>\n")
	} else {
		fmt.Printf("%v %T\n", v, v)
	}
}
