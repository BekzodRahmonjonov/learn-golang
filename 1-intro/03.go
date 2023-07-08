package main

import "fmt"

func main() {
	fmt.Println(fn(5, 6))
}

// 1
func fn(a int, b int) int {
	return a + b
}

func fn2(a, b int) int {
	return a + b
}
