package main

import "fmt"

// const a = 42
// const b = 42.78
// const c = "Test String"

// Two types of constants = typed and untyped
const (
	a int = 42
	b float64 = 42.78
	c string = "Test String"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
