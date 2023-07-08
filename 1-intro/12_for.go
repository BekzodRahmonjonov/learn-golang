package main

import "fmt"

func main() {
	baseFor()
	optinalInitAndPostStatements()
	forInsteadWhile()
	infiniteLoop()
}

func baseFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func optinalInitAndPostStatements() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forInsteadWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	for {
	}
}
