package main

import "fmt"

func main() {
	fmt.Println(fn(5, 6))
}

func fn(a, b int) int {
	return a + b
}
