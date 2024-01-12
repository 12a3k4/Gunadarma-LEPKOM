package main

import "fmt"

func main() {
	var kursus = []string{"dbms", "server_os", "networking", "web", "deskop", "erp"}
	kursus_saya := kursus[1:5]
	kursus_saya = append(kursus_saya, "Manajemen Informatika")

	fmt.Println("Isi dari kursus_saya adalah : ", kursus_saya)
}
