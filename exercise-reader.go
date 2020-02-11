package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
// Read 用数据填充给定的字节切片并返回填充的字节数和错误值
func (reader MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}


func main() {
	reader.Validate(MyReader{})

	a := 'a'
	fmt.Println(byte(a))
}


