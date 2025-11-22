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
func checkNumber(number string, check func(string) string) string { 
	//Function yg jadi parameter bisa dijadiin type data
	return fmt.Sprintf("Satuan dari %s adalah %s", number, check(number))
	
}

func satuan(number string) string {
	if len(number) == 1 {
		return "satuan"
	} else if len(number) == 2 {
		return "puluhan"
	} else if len(number) == 3 {
		return "ratusan"
	} else {
		return "ribuan atau lebih"
	}
}

// ================ RECURSIVE FUNCTION ================ //
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// ================ MAIN FUNCTION ================ //

func main() {
	name := sayName // Function as a value
	fmt.Println(name( "Alber"))

	fmt.Println(hewan("Kucing", "Anjing")) // Multiple return values


	fmt.Println(sumAll(1, 2, 3, 4, 5)) // Mengirim beberapa argumen ke varargs
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(sumAll(numbers...)) // Mengirim slice ke varargs

	filter := satuan
	fmt.Println(checkNumber("5", filter)) // Function as parameter
	fmt.Println(checkNumber("50", filter)) // Function as parameter
	fmt.Println(checkNumber("500", filter)) // Function as parameter
	fmt.Println(checkNumber("5000", filter)) // Function as parameter

	anonFunc := func(name string) string { // Anonymous function
		return "Hi " + name
	}
	fmt.Println(anonFunc("Alber")) // Memanggil anonymous function

	fmt.Println(factorial(5)) // Recursive function
	fmt.Println(5*4*3*2*1) // Pembuktian recursive function
}