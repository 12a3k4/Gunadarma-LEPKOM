package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " adalah angka genap")
		} else {
			fmt.Println(i, " adalah angka ganjil")
		}
	}
}
