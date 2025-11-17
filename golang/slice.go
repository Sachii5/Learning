package main

import "fmt"

func main()  {
	//  tipe data clice punya 3 data penting
	// 1. pointer : penunjuk data pertama di array para slice
	// 2. length : panjang dari slice
	// 3. capacity : kapasitas dari slice, dimana length tidak booleh lebih dari capacity

	hewan := [...]string{"kucing", "anjing", "ikan", "ayam"}

	slice1 := hewan[1:4]
	slice2 := hewan[:2]
	slice3 := hewan[3:]
	slice4 := hewan[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)


}