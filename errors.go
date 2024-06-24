package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64, n int) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z float64 = 1
	for i := 0; i < n; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	number := float64(-5)
	number_sqrt, err := Sqrt(number, 10)
	fmt.Printf("%v %v\n", number_sqrt, err)
}
