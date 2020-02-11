package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot sqrt negative number: %v", e) // 会陷入无限循环

	return fmt.Sprintf("cannot sqrt negative number: %v", float64(e))

}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	i := 0
	for i < 10 {
		z -= (z * z - x) / (2 * z)
		i++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


