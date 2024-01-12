package main

import "fmt"

func main() {
	x := 13
	for i := 1; i <= x; i++ {
		if i%2 == 1 {
			fmt.Println("Angka ", i)
		}
	}
}
