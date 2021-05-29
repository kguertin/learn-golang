package main

import "fmt"

var i = 43

// Sets type and sets zero value of that type if not assigned a value (nil, 0, false, etc.)
var z int

func main() {
	n, _ := fmt.Println("Hey!")
	fmt.Println(n)

	// Short declaration operator
	x := 42
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	y := 100 + 566
	fmt.Println(y)

	// Var can be used to create a variable outside function body, but should try to keep variables to limited scope
	fmt.Println(i)
	foo()

	fmt.Println(z)

}

func foo() {
	fmt.Println("Print Again:", i)
}
