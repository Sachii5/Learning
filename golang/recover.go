package main

import "fmt"

// recover adalah mekanisme untuk menangani panic agar program tidak berhenti secara tiba-tiba.
// dengan recover, kita bisa "memulihkan" eksekusi program setelah terjadi panic.

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover() // recover menangkap panic dan mengembalikan pesan errornya
	if message != nil {
		fmt.Println("Terjadi error:", message)
	}
}

func runApp(error bool) {
	defer endApp() // defer akan dieksekusi di akhir function runApp
	if error {
		panic("Aplikasi mengalami error") // panic akan menghentikan eksekusi program
	}
}	

func main() {
	runApp(true)
	fmt.Println("Aplikasi berjalan") // baris ini akan dieksekusi karena recover menangani panic
}