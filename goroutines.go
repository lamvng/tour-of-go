package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// If the Sleep() is commented. Output: Only "Main function" printed.
		// Because the main function exits before the coroutine can execute.
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// "Coroutine" and "Main function" are printed concurrently
func main() {
	go say("Goroutines")
	say("Main function")
}
