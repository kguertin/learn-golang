package main

import (
	"fmt"
	"runtime"
)


var x int
var y float64
var a int8 = 127
var b uint8 = 255

func main() {
	// Numeric Type (Integers and Floating Points)
	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	
}