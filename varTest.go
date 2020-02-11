package main

import (
	"fmt"
)

var c, python, java bool
var j, k int = 1, 2
var m, n, o = 3, "3", false


func main() {
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(j, k, m, n, o)

	// 用 := 代替var声明，但不能在函数外使用
	tt := true
	fmt.Println(tt)

}
