package main

import "fmt"

// 用牛顿法实现平方根函数

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for i < 10 {
		z -= (z * z - x) / (2 * z)
		i++
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
