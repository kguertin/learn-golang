package main

import "fmt"

func main() {
	num := 13
	fmt.Printf("%v\t%b\t%#x\n", num, num, num)
	newNum := num << 1
	fmt.Printf("%v\t%b\t%#x\n", newNum, newNum, newNum)

}