package main

import "fmt"

func main() {
	// tipe data map adalah tipe data yang berisi kumpulan data dengan format key dan value
	// key adalah penanda unik untuk setiap value
	// value adalah isi dari data yang disimpan di map

	// ======= MEMBUAT MAP ======= //
	nama := map[string]string{
	"depan": "Alber",
	"tengah": "Galih",
	"Belakang": "Anthoni"}
	fmt.Println(len(nama))

	fmt.Println(nama)
	fmt.Println(nama["depan"])
	fmt.Println(nama["tengah"])
	fmt.Println(nama["Belakang"])

	// ======= EDIT VALUE DI MAP ======= //
	nama["depan"] = "Khanabawi"
	fmt.Println(nama)
	fmt.Println(nama["depan"])

	//  ========= HAPUS DATA DI MAP ======== //
	delete(nama, "tengah")
	fmt.Println(nama)
}