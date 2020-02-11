package main

import "fmt"

func main() {
	// defer语句会将函数推迟到外层函数返回之后执行
	// 推迟调用的函数其“参数”会立即求值，但知道外层函数返回前该函数都不会被调用
	defer fmt.Println("world")
	fmt.Println("hello")
}
