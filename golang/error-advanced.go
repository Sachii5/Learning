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

type ProductRepo struct {
	products map[int]Product
}

type ProductService struct{
	repo *ProductRepo
}

var ErrProductNotFound = errors.New("product not found")
var ErrInvalidQuantity = errors.New("invalid quantity")
var ErrOutOfStock = errors.New("out of stock")

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

func (s *ProductService) Buy(id int, qty int) error {
	if qty <=0 {
		return ErrInvalidQuantity
	}
	
	p, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
		return ErrProductNotFound
		}
		return fmt.Errorf("Service buy : %w", err)
	}
	
	if p.Stock < qty {
		return ErrOutOfStock
	}
	
	NewStock := p.Stock - qty
	err = s.repo.UpdateStock(id, NewStock)
	if err != nil {
		return fmt.Errorf("Service buy : %w", err)
	}
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
	err = r.UpdateStock(1, 10)
	p, _ = r.FindByID(1)
	fmt.Println(p,err)
	
	s := ProductService{repo :&r}
	
	err = s.Buy(1, 5)
	p, _ = s.repo.FindByID(1)
	fmt.Println(p, err)
	
	err = s.Buy(2, 5)
	p, _ = s.repo.FindByID(2)
	fmt.Println(p, err)
	
	err = s.Buy(1, 11)
	p, _ = s.repo.FindByID(1)
	fmt.Println(p, err)
}