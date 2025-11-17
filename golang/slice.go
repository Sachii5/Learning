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

	// ======== MENGUBAH ISI ARRAY DENGAN SLICE ====== //

	sliceHewan := hewan[2:]
	sliceHewan[0] = "patin"
	fmt.Println(hewan)

	// ==================== APPEND() ================== //
	
	sliceHewan2 := append(sliceHewan, "rusa")
	sliceHewan2[1] = "jago"
	fmt.Println(sliceHewan2)
	fmt.Println(hewan)

	// ================== MAKE() ==================== //

	newSlice := make([]string, 3, 6)
	newSlice[0] = "Alber"
	newSlice[1] = "Alber"
	newSlice[2] = "Alber"

	fmt.Println(newSlice)

	newSlice2 := append(newSlice, "Alber")
	newSlice[0] = "Galih"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// ================= COPY() =================== //

	allHewan := hewan[:]
	makhlukHidup := make([]string, len(allHewan), cap(allHewan))

	copy(makhlukHidup, allHewan)

	fmt.Println(makhlukHidup)
	fmt.Println(len(makhlukHidup))
	fmt.Println(cap(makhlukHidup))

}