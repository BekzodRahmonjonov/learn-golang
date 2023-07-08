package main

import "fmt"

/*
a := "Bekzod" - It's wrong!!! Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/
func main() {
	var i, j int = 1, 2
	k := 3 // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	c, python, javascript := true, false, "no!"

	fmt.Println(i, j, k, c, python, javascript) // 1, 2, 3, true, false, no!
}
