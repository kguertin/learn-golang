package main

import "fmt"

var a int
var b string 
var c bool

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v %v %v\n",x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Values of package scope variables
	// These have been assigned zero values 
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

}