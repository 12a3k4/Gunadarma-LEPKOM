package main

import "fmt"

func main() {
	x := 5
	for i := 0; i < x; i++ {
		for x := 4; i <= x; x-- {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
