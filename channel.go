package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)     // Make a channel
	go sum(s[:len(s)/2], c) // -5
	go sum(s[len(s)/2:], c) // 17
	x, y := <-c, <-c        // Receive from c
	fmt.Println(x, y, x+y)  // -5 17 12 // FIFO output
}
