package main

import "fmt"

var i = 43

// Sets type and sets zero value of that type if not assigned a value (nil, 0, false, etc.)
// Package scope
var z int
var str string = "This is a string"
var a string =  `He said "NO!"`

// This will infer the type based on the assigned value
var b = 13

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

	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", str)
	fmt.Println(a)

	fmt.Println(y)
	// Format Strings as first value
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)
	fmt.Printf("%v\n", b)

	//String Print
	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)

    // Create Types
	c := 666
	fmt.Printf("%T\n", c)
    
	type hotdog int
	var d hotdog = 21

	fmt.Printf("%T\n", d)

}

func foo() {
	fmt.Println("Print Again:", i)
}
