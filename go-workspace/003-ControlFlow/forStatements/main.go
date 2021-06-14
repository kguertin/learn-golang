package main

import "fmt"

func main() {
	x := 1

	for x < 10 {
		fmt.Printf("x: %d\n", x)
		x++
	}

	y := 1
	for {
		if y > 10 {
			break
		}
		fmt.Printf("y: %d\n", y)
		y++
	}

    z := 0
	for {
		z++
		if z > 100{
			break
		}

		if z % 2 != 0{
			continue
		}
		fmt.Printf("z: %d\n", z)
	}
	
}

