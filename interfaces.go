package main

import (
	"fmt"
	"math"
)

// 接口类型是由一组方法签名定义的集合
type Abser interface {
	Abs() float64
}

type JoJoFloat float64

func (f JoJoFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser // 保存任何实现了Abser定义的方法集合的值
	f := JoJoFloat(-math.Sqrt2)
	v := Vertex8{3, 4}
	a = f
	fmt.Println(a.Abs())
	a = &v // Abs方法只为*Vertex（指针类型）定义
	fmt.Println(a.Abs())
}
