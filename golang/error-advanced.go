package main

import (
	"fmt"
	"errors"
)

type Product struct {
	ID int
	Name string
	Stock int
}

var ErrProductNotFound = errors.New("product not found")
var ErrInvalidQuantity = errors.New("invalid quantity")
var ErrOutOfStock = errors.New("out of stock")

type ProductRepo struct {
	products map[int]Product
}

func (r *ProductRepo) FindByID(id int) (Product, error) {
	P, ok := r.products[id]
	if !ok{
		return Product{}, ErrProductNotFound
	}
	return P, nil
}

func (r *ProductRepo) UpdateStock(id int, update int) error{
	P, ok := r.products[id]
	if !ok{
		return ErrProductNotFound
	}

	P.Stock = update 
	r.products[id] = P
	return nil
}

func main() {
	r := ProductRepo{
		products: map[int]Product {
			1: {1, "Mie", 5},
		},
	}

	p, err := r.FindByID(1)
	fmt.Println(p,err)
	err = r.UpdateStock(1, 2)
	p, _ = r.FindByID(1)
	fmt.Println(p,err)
}