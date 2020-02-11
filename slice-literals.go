package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13} // 创建一个数组，然后构建了一个引用它的切片
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
