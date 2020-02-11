package main

import "fmt"

func main() {
	// defer语句后的函数调用会被压入到一个栈中
	// 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	fmt.Println("counting")
	for i:=0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
