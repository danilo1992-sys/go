package main

import "fmt"

func main() {
	nombre, edad := retornomultiple()
	fmt.Printf("Mi %v ,edad es %v \n", nombre, edad)
}

func retornomultiple() (string, int) {
	return "Danilo", 33
}
