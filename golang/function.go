package main

import "fmt"

// ======== FUNCTION AS A VALUE ========== //

func sayName(name string)string {
	name = "Hello " + name
	return name
}

// ======== MULTIPLE RETURN VALUES ========== //

func hewan(hewan1 string, hewan2 string) (string, string){
	return hewan1, hewan2
}

// ========= VARARGS / VARIABLE ARGUMENTS ========== //

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// ================= FUNCTION AS PARAMETER ================ //
func checkNumber(number int, check func(int) string) string {
	fmt.Println("Satuan dari", number, "adalah" , check(number))
}

func satuan(number int) string {
	if len()
}

// ================ MAIN FUNCTION ================ //

func main() {
	name := sayName // Function as a value
	fmt.Println(name("Alber"))

	fmt.Println(hewan("Kucing", "Anjing")) // Multiple return values


	fmt.Println(sumAll(1, 2, 3, 4, 5)) // Mengirim beberapa argumen ke varargs
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(sumAll(numbers...)) // Mengirim slice ke varargs
}