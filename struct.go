package main

import "fmt"

// Struct (Similar to C)
type Vertex struct {
	X int
	Y int
}

func main() {

	// Struct
	vertex := Vertex{1, 2}
	vertex.X = 4
	p := &vertex
	fmt.Println(vertex.X)
	fmt.Println(p.Y)

}
