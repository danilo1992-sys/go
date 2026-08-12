package main

import (
	"fmt"
	"strings"
)

func main() {
	cadena := "hola mundo"
	fmt.Println(strings.ToUpper(cadena))
	fmt.Println(strings.ToLower(cadena))
	letras := strings.Split(cadena, "")
	fmt.Println(letras)
	pos := strings.Index(cadena, "hola")
	if pos == -1 {
		fmt.Println("La palabra no esta dentro del texto", cadena)
	} else {
		fmt.Println("La palabra esta dentro del texto", cadena, "en la posicion", pos)
	}

	repetidas := strings.Repeat(cadena, 10)
	fmt.Println(repetidas)

	cadena2 := strings.Replace(cadena, "hola", "nuevo", -1)
	fmt.Println(cadena2)

	fmt.Println(string(cadena[0:2]))
}
