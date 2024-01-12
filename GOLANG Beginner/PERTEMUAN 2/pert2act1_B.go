package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f_name string = "Ali"
	I_name := "Akbar"
	age := 20

	// Output 1
	// fmt.Printf("My name is %s %s, Iam %d old\n", f_name, I_name, age)

	// Output 2
	fmt.Print("My name is ", f_name, " ", ", I_name, ", " Iam ", strconv.Itoa(age), " old")

}
