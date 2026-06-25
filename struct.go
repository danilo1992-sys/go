package main

import "fmt"

func main() {
	estructura := Personas{
		Id:     1,
		Nombre: "Danilo Caceres",
		Email:  "test@test.com",
		Edad:   33,
	}
	fmt.Println(estructura)

	p := new(estructura)
	p.Id=2
	p.Nombre="juan"
	p.Edad=22
	p.Email="hola@hola.com"

	fmt.Println(p)
}

type Personas struct {
	Id     int
	Nombre string
	Email  string
	Edad   int
}
