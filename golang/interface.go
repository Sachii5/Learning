package main 

import "fmt"

type Makhluk interface {
	suara(bunyi string) string
}

type Manusia struct {}

type Hewan struct {}

func (m Manusia) suara(bunyi string) string {
	return bunyi
}

func main() {
	var m Makhluk

	m = Manusia{}
	fmt.Println(m.suara("Halo, saya manusia"))
}