package main

import "fmt"

func main() {
	var i interface{} // 空接口

	// 空接口可保存任何类型的值，因为每个类型都至少实现了零个方法
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)
	i = nil
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
