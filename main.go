package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {

	a := 1
	b := 2

	rs := sum(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, rs)
}
