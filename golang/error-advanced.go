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
	id, ok := r.products[id]
	if !ok{
		return Product{}, ErrProductNotFound
	}
	return user, nil
}

func (r *ProductRepo) UpdateStock(id int, update int) error{
}

func main() {
	Produk := Product{1, "mie", 5}
	Repo := ProductRepo{}

}