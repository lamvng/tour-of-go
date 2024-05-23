package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	point := Vertex{1, 2}
	point.X = 4
	fmt.Println(point)
}
