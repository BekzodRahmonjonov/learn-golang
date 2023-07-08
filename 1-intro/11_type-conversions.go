package main

import "fmt"

// The expression T(v) converts the value v to the type T.
// Some numeric conversions:
var o int = 42
var f float64 = float64(o)
var u uint = uint(f)

func main() {
	// Or, put more simply:
	o2 := 42
	f2 := float64(i)
	u2 := uint(f)

	fmt.Println(o, f, u, o2, f2, u2)
}
