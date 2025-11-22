package main

import "fmt"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	// defer akan mengeksekusi sebuah function di akhir sebuah function lain
	defer fmt.Println(sayGoodbye("Alber")) // Defer akan dieksekusi di akhir function main
	fmt.Println("Hello Alber")
}