package main

import "fmt"

func main() {
	Hola()
}

func Hola() {
	defer fmt.Println("mensaje con defer")
	fmt.Println("Primer mensaje")
	a := 1
	if a == 1 {
		panic("Fallo con exito")
	}
}
