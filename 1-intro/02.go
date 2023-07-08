package main

import (
	"fmt"
	"strconv"
)

func main() {
	// result := int(float64(42) + math.Pi)

	a := "12a"
	atoi, _ := strconv.Atoi(a)

	fmt.Println(atoi)

	//fmt.Println(result)
}

func myFunc() string {
	return ""
}
