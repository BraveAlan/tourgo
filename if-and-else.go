package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// %g  根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// 这里开始就不能使用变量v了
	// fmt.Println(v) // 代码检查直接会报错
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}