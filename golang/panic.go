package main

import "fmt"

// panic adalah mekanisme untuk menghentikan eksekusi program secara tiba-tiba ketika terjadi error yang tidak terduga.
// ketika panic terjadi, program akan menghentikan eksekusi dan menampilkan pesan error beserta stack trace.

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp() // defer akan dieksekusi di akhir function runApp
	if error {
		panic("Aplikasi mengalami error") // panic akan menghentikan eksekusi program
	}
}

func main() {
	runApp(true)
	fmt.Println("Aplikasi berjalan") // baris ini tidak akan dieksekusi karena panic terjadi sebelumnya
}