package main

import "fmt"

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2} // 创建一个Vertex2类型的结构体
	v2 = Vertex2{X: 1} // Y:0被隐式地赋予
	v3 = Vertex2{} // X:0 Y:0
	p = &Vertex2{1, 2} //创建一个*Vertex2类型的结构体指针

)

func main() {
	fmt.Println(v1,p,v2,v3)
}
