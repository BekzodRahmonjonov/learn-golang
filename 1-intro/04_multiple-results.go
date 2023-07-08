package main

import "fmt"

func swap(x, y string) (string, string) {
	return x, y
}

// Syntax sugar

func swap1(x, y string) (a, b string) {
	a = x
	b = y
	return a, b
}

func main() {
	a, b := swap("Bekzod", "Oybek")
	fmt.Println(a, b)
}
