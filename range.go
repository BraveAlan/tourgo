package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// i为下标，v为该下标所对应元素的一份副本
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
