package main

import "fmt"

func main() {
	var nama [3]string

	nama[0] = "Alber"
	nama[1] = "Galih"
	nama[2] = "Anthoni"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	fmt.Println(len(nama))

}
