package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum * 2
	return // this function return us x and y, it's also syntax sugar. We can write this like this also - return x, y
}

func main() {
	fmt.Println(split(10)) // 5, 20
}
