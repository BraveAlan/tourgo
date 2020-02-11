package main

// 使用指针接收者的原因：
//  1. 方法能够修改其接收者指向的值
//  2. 避免在每次调用方法时复制该值，若值的类型为大型结构体时，这样做会更加高效

import (
	"fmt"
	"math"
)

type Vertex7 struct {
	X, Y float64
}

func (v *Vertex7) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex7{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
