package main

import "fmt"

func main() {
	primes := [6]int{2,3,5,7,11,13}
	// 创建一个切片，包含a中下标从1到3的元素
	var s []int = primes[1:4] // 半开区间，包括第一个元素，排除最后一个元素
	fmt.Println(s)
}
