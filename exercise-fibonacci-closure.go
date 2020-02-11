package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int{
	cur := 0
	next := 1
	return func() int {
		result := cur
		cur = next
		next = result + next
		return result
	}
}

func main() {
	f := fibonacci()
	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}
