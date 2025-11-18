package main

import "fmt"

func sayName(name string)string {
	name = "Hello " + name
	return name
}

func hewan(hewan1 string, hewan2 string) (string, string){
	return hewan1, hewan2
}

func main() {
	name := sayName("Alber")
	fmt.Println(name)

	fmt.Println(hewan("Kucing", "Anjing"))
}