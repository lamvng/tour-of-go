// https://go.dev/tour/flowcontrol/8

package main

import (
	"fmt"
)

func sqrt(x float64, n int) float64 {
	var z float64 = 1
	for i := 0; i < n; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Loop %d: z = %f\n", i, z)
	}
	return z
}

func main() {
	sqrt(2, 10)
}
