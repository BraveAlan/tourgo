package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 使用指针接收者的方法可以修改接收者指向的值
func (v *Vertex5) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex5{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
