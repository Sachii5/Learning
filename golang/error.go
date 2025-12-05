package main

import (
	"fmt"
	"errors"
)

var errInvalidQty = errors.New("invalid quantity")

func OrderProduct(qty int) error {
	if qty <= 0 {
		return errInvalidQty
	}
	return nil
}

func main() {
	err := OrderProduct(0)
	if err != nil {
		fmt.Println("Error ordering product:", err)
	} else {
		fmt.Println("Product ordered successfully")
	}

}