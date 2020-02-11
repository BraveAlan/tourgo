package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	names2 := [2]string{"JOJO", "DIO"}
	fmt.Println(names)
	fmt.Println(names2)

	// 切片就像数组的引用，理解为指针，更改切片的元素会修改其底层数组中对应的元素
	// 当元素发生变化时，与它共享底层数组的切片都会发生变化
	a := names[:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "JOJO" // b[0]映射到原数组中就是names[1]，该元素是a和b的公共元素
	fmt.Println(a, b)
	fmt.Println(names)
}
