// https://code101.medium.com/understanding-closures-in-go-encapsulating-state-and-behaviour-558ac3617671

package main

import "fmt"

func createCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment // createCounter() return another function
}

func main() {
	counter1 := createCounter()
	counter2 := createCounter()

	// Function variables are saved

	fmt.Println(counter1()) // 1 // States are maintained
	fmt.Println(counter1()) // 2
	fmt.Println(counter2()) // 1
	fmt.Println(counter2()) // 2
}
