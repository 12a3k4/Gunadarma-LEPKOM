package main

import "fmt"

func naikanGaji(gajiAwal int, gajiAkhir int) int {
	gajiAwal = gajiAkhir
	return gajiAwal
}

func main() {
	var gajiSekarang, ekspektasigaji int

	fmt.Print("Masukkan Gaji Anda : ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan Gaji yang Anda inginkan : ")
	fmt.Scan(&ekspektasigaji)

	naikanGaji(gajiSekarang, ekspektasigaji)

	fmt.Printf("\nGaji Anda Sekarang %d\n", naikanGaji(gajiSekarang, ekspektasigaji))
}
