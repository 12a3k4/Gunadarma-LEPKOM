package main

import "fmt"

func main() {
	var l int
	for i := 0; i < 10; i++ {
		fmt.Print("Input Nilai l :")
		fmt.Scan(&l)

		if l > 10 {
			fmt.Println("Pertanyaan selesai, karena angka 1 yang diinput sudah melebihi dari 10. Terima Kasih")
			break
		} else if l%2 == 1 {
			fmt.Println(l, " adalah Angka Ganjil")
		} else {
			fmt.Println(l, " adalah Angka Genap")
		}
	}
}
