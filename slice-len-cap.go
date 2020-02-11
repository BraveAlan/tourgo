package main

import "fmt"

func main() {
	// 切片的长度len和容量cap
	// 切片的长度就是切片所包含的元素个数
	// 切片的容量就是从切片的第一个元素到其底层数组元素末尾的个数
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 继续拓展其长度
	s = s[:6]
	printSlice(s)

	// 继续拓展其长度，会报错 slice bounds out of range
	//s = s[:10]
	//printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
