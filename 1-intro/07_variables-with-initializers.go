package main

import "fmt"

var i, j int = 1, 2 // there i equal to 1 and j equal to 2. Type of variables is integer

func main() {
	var c, python, javascript = true, false, "no!" // If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

	fmt.Println(i, j, c, python, javascript) // 1, 2, true, false, no!
}
