package main

import "fmt"

/*
Pointer adalah variabel yang menyimpan alamat memori dari variabel lain.
Dengan pointer, kita dapat mengakses dan memodifikasi nilai variabel tersebut secara langsung melalui alamat memorinya.
tanpa pointer, kita hanya mengoper salinan nilai variabel ke dalam function.
jadi jika kita mengubah nilai di dalam function, nilai asli di luar function tidak akan berubah.
dengan pointer, kita mengoper alamat memori variabel ke dalam function.
sehingga jika kita mengubah nilai di dalam function, nilai asli di luar function juga akan berubah.
*/

// fungsi tambah tanpa pointer
func tambah(b int) {
	b = b + 1
}

// fungsi tambah dengan pointer
func tambahPointer(a *int) {
	*a = *a + 1
}

func main() {
	// tanpa pointer
	b:= 10
	tambah(b)
	fmt.Println(b) // Output: 10
	// dengan pointer
	a:= 10
	tambahPointer(&a)
	fmt.Println(a) // Output: 11
}