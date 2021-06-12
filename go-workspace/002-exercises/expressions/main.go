package main

import "fmt"

const (
	num1     = 13
	num2 int = 666
)

func main() {
	a := 5 == 5
	fmt.Println(a)
	b := 3 <= 5
	fmt.Println(b)
	c := 3 >= 5
	fmt.Println(c)
	d := 5 != 5
	fmt.Println(d)
	e := 5 < 10
	fmt.Println(e)
	f := 5 > 10
	fmt.Println(f)

	fmt.Printf("%v\t%T\n", num1, num1)
	fmt.Printf("%v\t%T\n", num2, num2)

}
