package main

import "fmt"

var x bool

func main() {
	
	// Zero value of bool is flase
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	// Will evaluate this expression
	// No === like js
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println( a > b)
}