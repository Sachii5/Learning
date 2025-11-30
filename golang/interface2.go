package main

import "fmt"

type Wallet interface{
	GetBalanced() int
	Add(amount int)
}

type Ewallet struct{
	Balanced int
}

func (e Ewallet) GetBalanced() int{
	return e.Balanced
}

func (e Ewallet) Add(amount int) {
	e.Balanced += amount
}

func TopUp(e Wallet, amount int){
	e.Add(amount)
	fmt.Println("Succed")
}

func main()  {
	e := Ewallet{Balanced: 10000}
	TopUp(e, 15000)
}