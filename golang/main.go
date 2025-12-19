package main

import (
	"encoding/json"
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

type BuyRequest struct{
	ID int `json:"id"`
	Qty int `json:"qty"`
}

type SuccessResponse struct{
	Status string `json:"status"`
	Message string `json:"message"`
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

func (s *ProductService) GetProduct(id int) (Product, error){
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
	
func (s *ProductService) Buy(id int, qtyBuy int) error{
	if id <= 0 {
		return ErrInvalidId
	}
	
	if qtyBuy <= 0 {
		return ErrInvalidQuantity
	}
	
	prod, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			return ErrProductNotFound
		}
		return fmt.Errorf("Service buy : %w", err)
	}
	
	if prod.Stock < qtyBuy {
		return ErrOutOfStock
	}
	
	newQty := prod.Stock - qtyBuy
	err = s.repo.UpdateStock(id,newQty)
	if err != nil{
		return fmt.Errorf("Service buy : %w", err)
	}
	
	return nil
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
	
	product, err := s.GetProduct(id)
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

func (s *ProductService) buyHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}
	
	var req BuyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err !=nil {
		http.Error(w, "Bad request", 400)
		return
	}
	
	if req.ID <= 0 {
		http.Error(w, "Invalid ID", 400)
		return
	}
	
	if req.Qty <= 0 {
		http.Error(w, "Invalid QTY", 400)
		return
	}
	
	errBuy := s.Buy(req.ID, req.Qty)
	if errBuy != nil {
		switch {
			case errors.Is(errBuy, ErrInvalidId):
				http.Error(w, "Invalid ID", 400)
			case errors.Is(errBuy, ErrInvalidQuantity):
				http.Error(w, "Invalid QTY", 400)
			case errors.Is(errBuy, ErrProductNotFound):
				http.Error(w, "Prodcut not found", 404)
			case errors.Is(errBuy, ErrOutOfStock):
				http.Error(w, "Out of stock", 409)
			default:
				http.Error(w, "Internal server error", 500)
		}
		return
	} 	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(SuccessResponse{
		Status: "Success",
		Message: "Buy success",
	})
}

func main() {
	r := ProductRepo{
		products: map[int]Product {
			1: {1, "Mie", 5},
		},
	}
	s := ProductService{repo :&r}
	
	http.HandleFunc("/product", s.productHandler)
	http.HandleFunc("/buy", s.buyHandler)
	
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	

	http.ListenAndServe(":8080", nil)

}