package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m map[string]Vertex
	m = make(map[string]Vertex) // make map[key-type]value-type

	// Add element
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Google"] = Vertex{37.42202, -122.08408}

	// Test if key is present
	elem, is_elem_exist := m["Google"]
	fmt.Printf("%t %v\n", is_elem_exist, elem)

	elem_non_exist, is_elem_non_exist := m["Facebook"]
	fmt.Printf("%t %v\n", is_elem_non_exist, elem_non_exist) // elem_non_exist = {0 0}
}
