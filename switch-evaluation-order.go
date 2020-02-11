package main

import (
	"fmt"
	"time"
)

func main() {
	// switch case 的求值顺序是从上到下顺次执行，直到匹配成功时停止
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	fmt.Println("today:", today)
	fmt.Println("Saturday:", time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}
