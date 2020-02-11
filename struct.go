package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4 // 使用点号访问字段
	fmt.Println(v.X)
	// 结构体指针
	p := &v
	p.X = 1e9 // 隐式间接引用，即(*p).X
	fmt.Println(v)

}
