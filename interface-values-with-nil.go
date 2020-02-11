package main

import "fmt"

type I3 interface {
	M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I3
	var t *T3
	i = t     // 保存了 nil 具体值的接口其自身并不为 nil
	describe(i)
	i.M()
	if i == nil { // i本身不为nil，它保存了T3类型值
		fmt.Println("it is nil")
	}

	i = &T3{"hello"}
	describe(i)
	i.M()
}

func describe(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}