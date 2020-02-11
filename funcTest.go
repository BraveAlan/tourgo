package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func returnEmpty(test int) {
	fmt.Println(test)
}

func main() {
	a, b := swap("1", "2")
	fmt.Println(a, b)
	fmt.Println(split(17))
	returnEmpty(0)
}

