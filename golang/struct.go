package main

import "fmt"

type User struct {
	Name     string
	Address  string
	Age      int
}

func TambahUser(user *User) User {
	user.Name = "Galih"
	user.Address = "Jepang"
	user.Age = 20
	return *user
}

// Struct method
func (user User) sayHello()  {
	fmt.Println("Halo ", user.Name)
}

func main()  {
	user := User{"Alber", "Indonesia", 15}
	TambahUser(&user)
	fmt.Println(user)
	fmt.Println("Name:", user.Name)
	fmt.Println("Address:", user.Address)
	fmt.Println("Age:", user.Age)

	fmt.Println("===================================")
	
	user2 := User{Name : "Alber"}
	user2.sayHello() // Struct method
	
}