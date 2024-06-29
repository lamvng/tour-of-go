package main

import "fmt"

func main() {
	lengthChannel := 4
	ch := make(chan int, lengthChannel)
	// for i := range lengthChannel {
	for i := range lengthChannel + 1 {
		ch <- i // fatal error: all goroutines are asleep - deadlock!
	}

	for _ = range ch {
		fmt.Println(<-ch)
	}
}
