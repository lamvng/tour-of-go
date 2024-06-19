package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f_plus_1, f_plus_2 := 0, 1
	return func() int {
		f_current := f_plus_1
		f_plus_1, f_plus_2 = f_plus_2, f_plus_1+f_plus_2
		fmt.Printf("f_current = %4d f_plus_1 = %4d f_plus_2 = %4d\n", f_current, f_plus_1, f_plus_2)
		return f_current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		f()
	}
}
