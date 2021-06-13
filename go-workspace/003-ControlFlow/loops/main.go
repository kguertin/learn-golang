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
}