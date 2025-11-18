package main

import "fmt"

func main()  {
	if length := len("Anthoni"); length > 5 {
		fmt.Println("Panjang string lebih dari 5 karakter:", length)
	} else {
		fmt.Println("Panjang string 5 karakter atau kurang:", length)
	}
}