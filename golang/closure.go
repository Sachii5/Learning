package main

import "fmt"


// closuse adalah kemampuan function untuk "mengakses" variable di luar lingkup function tersebut.
// hati-hati dalam menggunakan closure karena bisa menyebabkan efek samping (side effect) 
// pada variable di luar lingkup function

func main() {
	name := "Alber"
	fmt.Println(name)

	// anonymous function
	sayHello := func() {
		name = "Galih"
	}
	sayHello()
	fmt.Println(name)
}