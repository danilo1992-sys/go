package main

import "fmt"

func main() {
	Tabla := tabla(5)
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %v = %v \n", i, Tabla())
	}
}

func tabla(valor int) func() int {
	numero := valor
	sec := 0
	return func() int {
		sec++
		return numero * sec
	}
}
