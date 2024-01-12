package main

import "fmt"

func main() {

	var data = map[string]mahasiswa{
		"50421119": {
			"Ali Akbar Said",
			"2IA03",
		},
		"12131416": {
			"Mega Wati",
			"4DI41",
		},
	}
	var search string
	fmt.Print("Masukkan Npm Anda ? ")
	fmt.Scanf("%s", &search)

	var value, ok = data[search]

	if ok {
		fmt.Printf("Nama anda %s \nkelas anda %s", value.Nama, value.kelas)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}

type mahasiswa struct {
	Nama  string
	kelas string
}
