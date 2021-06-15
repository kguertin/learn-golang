package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	if 2 != 2 {
		fmt.Println("006")
	}

	if !(2 == 2) {
		fmt.Println("007")
	}

	if !(2 != 2) {
		fmt.Println("008")
	}

	// Can define x here and scope of x is within the if statement
	if x := 42; x != 2 {
		fmt.Println(x)
	}

	num := 13

	if num == 12 {
		fmt.Println("if")
	} else if num == 13 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
}
