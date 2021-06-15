package main

import "fmt"

func main() {
	// for init; condtion; post;{
	// }

	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("The outer loop: %v\t The inner loop: %d\n", i, j)
		}
	}

	for i := 33; i <= 122; i++ {
		fmt.Println(i)
		fmt.Printf("%#x\t%#U\n", i,i)
	}
}
