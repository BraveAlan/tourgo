package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	visit(t, ch)
	close(ch)
}

func visit(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		visit(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		visit(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i:=0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		fmt.Println("v1", v1)
		fmt.Println("v2", v2)
		if v1 !=v2 {
			return false
		}
	}
	return true
}

func main() {
	result := Same(tree.New(1), tree.New(1))
	fmt.Println(result)
	result = Same(tree.New(1), tree.New(2))
	fmt.Println(result)
}
