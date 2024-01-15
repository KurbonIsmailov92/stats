package main

import "fmt"

func main() {
	a := 10
	b := 15
	fmt.Println(sum(a, b))
}

func sum(a, b int) int {
	var c int = a + b
	return c
}
