package main

import "fmt"

func main() {
	var a [2]string
	var b [2]bool
	var c [2]int
	var d [2]float64
	fmt.Println(a[0])
	fmt.Println(b[0])
	fmt.Println(c[0])
	fmt.Println(d[0])

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
	// 数组长度是其类型的一部分，因此数组大小是不可变的
	//primes[6] = 17
	//fmt.Println(primes)
}
