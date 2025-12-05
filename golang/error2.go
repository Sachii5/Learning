package main
import (
	"fmt"
	"errors"
)

var errInvalidAge = errors.New("invalid age")

func isAdult(age int) (bool, error) {
	if age <= 0 {
		return false, errInvalidAge
	} else if age >= 18 {
		return true, nil
	} else {
		return false, nil
}
}

func main() {
	age, err := isAdult(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else if !age {
		fmt.Println("You are not an adult.")
	} else {
		fmt.Println("You are an adult.")
	}
}