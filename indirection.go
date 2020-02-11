package main

// 以指针为接收者的方法被调用时，接收者既能为值又能为指针
// 以值为接收者的方法被调用时，接收者既能为值又能为指针
import "fmt"

type Vertex6 struct {
	X, Y float64
}

func (v *Vertex6) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex6, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex6) Add() float64 {
	return v.X + v.Y
}

func main() {
	v := Vertex6{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println(v.Add())

	p := &Vertex6{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(p.Add())

	fmt.Println(v, p)
}
