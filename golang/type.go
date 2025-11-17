package main

import "fmt"

func main()  {
	type user string
	var users user

	users = user("Alber")

	fmt.Println(users)
}