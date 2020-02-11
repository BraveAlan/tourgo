package main

import "fmt"

func main() {
	// nil切片的长度和容量为0且“没有底层数组”
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
