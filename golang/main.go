package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
var ErrInvalidId = errors.New("Invalid ID")

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

func (s *ProductService) Product(id int) (Product, error){
	if id <= 0 {
		return Product{}, ErrInvalidId
	}
	
	p, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, ErrProductNotFound){
			return Product{}, ErrProductNotFound
		}
		return Product{}, fmt.Errorf("Service product : %w", err)
	}
	return p, nil
	}
	

func (s *ProductService) productHandler(w http.ResponseWriter, r *http.Request){
	idstr := r.URL.Query().Get("id")
	if idstr == ""{
		http.Error(w, "Missing id", 400)
		return
	}
	
	id, err := strconv.Atoi(idstr)
	if err != nil{
		http.Error(w, "Invalid id", 400)
		return
	}
	
	product, err := s.Product(id)
	if err != nil{
		switch{
		case errors.Is(err, ErrInvalidId):
			http.Error(w, "Invalid ID", 400)
		case errors.Is(err, ErrProductNotFound):
			http.Error(w, "Product not found", 404)
		default:
			http.Error(w, "Internal server error", 500)
		}
		return
	}
	fmt.Fprintln(w, "Product id :", product)
}

func main() {
	r := ProductRepo{
		products: map[int]Product {
			1: {1, "Mie", 5},
		},
	}
	s := ProductService{repo :&r}
	
	http.HandleFunc("/product", s.productHandler)
	
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	

	http.ListenAndServe(":8080", nil)

}