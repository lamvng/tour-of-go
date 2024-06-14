package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%3d cap=%3d %v\n", len(s), cap(s), s)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Another way of declaration
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var slice_prime []int = primes[0:4]
	printSlice(slice_prime)

	// Change value of slice
	slice_prime[3] = 17
	printSlice(slice_prime)
	fmt.Println(primes)

	// Slice literals
	sequence_slice := []int{1, 2, 3, 4}
	printSlice(sequence_slice)

	// Make a slice
	slice_make := make([]int, 5)
	printSlice(slice_make)

	slice_make_with_capacity := make([]int, 4, 5)
	printSlice(slice_make_with_capacity)

	// Append
	slice_make_with_capacity = append(slice_make_with_capacity, 1, 2, 3)
	printSlice(slice_make_with_capacity)

	// Range: Iterate over a slice
	// Return index and value
	// Skip return by _
	for index, value := range slice_make_with_capacity {
		fmt.Printf("Elem %3d = %3d\n", index, value)
	}
	for _, value := range slice_make_with_capacity {
		fmt.Printf("%d\n", value)
	}
}
