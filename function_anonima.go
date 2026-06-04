package main

import "fmt"

func main() {
	fmt.Print("La suma es ", suma(10, 12))
}

var suma = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}
