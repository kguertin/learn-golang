package main

import "fmt"

var a int
var b string
var c bool

var d int = 42
var e string = "James Bond"
var f bool = true

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v %v %v\n", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Values of package scope variables
	// These have been assigned zero values
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	s := fmt.Sprintf("%v\t%v\t%v", d, e, f)
	fmt.Println(s)

	type newType int
	var newVar newType = 5
	fmt.Printf("%v\n", newVar)
	fmt.Printf("%T\n", newVar)
	newVar = 42
	fmt.Printf("%v\n", newVar)

}
