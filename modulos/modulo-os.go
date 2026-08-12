package main

import (
	"flag"
	"fmt"
)

func main() {
	nombre := flag.String("nombre", "", "El nombre es")
	edad := flag.Int("edad", 18, "La edad es")
	flag.Parse()
	fmt.Println("El nombre es", *nombre)
	fmt.Println("La edad es", *edad)
}
