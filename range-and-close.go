package main

import "fmt"

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for i := range c { // 不添加close，range会报错，会报死锁，因为其他goroutine已执行完，永远不会有数据放入管道中
		fmt.Println(i)
	}
}
