package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 当缓冲区填满时，向其发送数据才会阻塞
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
