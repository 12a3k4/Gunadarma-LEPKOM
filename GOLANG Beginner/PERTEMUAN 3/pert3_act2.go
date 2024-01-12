package main

import "fmt"

func main() {
	var uts, uas int
	fmt.Print("Masukkan Nilai UTS : ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan Nilai UAS : ")
	fmt.Scan(&uas)

	total := (uts * 30 / 100) + (uas * 70 / 100)
	fmt.Println("Nilai UTS Anda : ", uts)
	fmt.Println("Nilai UAS Anda : ", uas)
	fmt.Println("Total Nilai Anda : ", total)

	if total >= 81 && total <= 100 {
		fmt.Println("Grade A")
	} else if total >= 61 && total <= 80 {
		fmt.Println("Grade B")
	} else if total >= 41 && total <= 60 {
		fmt.Println("Grade C")
	} else if total >= 21 && total <= 40 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade E")
	}
}
