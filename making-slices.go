package main

import "fmt"

func main() {
	// make函数会分配一个元素为0值得数组并返回一个引用了它的切片
	a := make([]int, 5)
	printSlice("a", a)

	// a等价于aa
	aa := []int{0,0,0,0,0}
	printSlice("aa", aa)

	b := make([]int, 0, 5) // 指定长度为0，容量为5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}