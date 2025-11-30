package main

import (
	"fmt"
)

type User struct{
	ID, Balanced int
	Name string 
}


func (u User) Display() {
	fmt.Println("ID", u.ID)
	fmt.Println("Name", u.Name)
	fmt.Println("Balanced", u.Balanced)
}

func (u User) Deposit(d int){
	u.Balanced += d
	fmt.Println("You deposited ", d, ", Now your balanced is ", u.Balanced)
}

func main()  {
	u := User{ID : 1, Name : "Alber", Balanced : 10000}
	u.Display()
	u.Deposit(15000)
}