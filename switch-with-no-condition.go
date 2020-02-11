package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // 等价于 switch true，即下面的case为true则匹配成功
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
