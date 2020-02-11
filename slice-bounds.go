package main

import "fmt"

func main() {
	/* 切片的默认行为
	对于数组
		var a[10] int
	来说，以下切片是等价的：
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}

