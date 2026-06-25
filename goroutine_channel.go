package main

import (
	"fmt"
	"time"
)

func main() {
	// ejemplo 1
	fmt.Println(retorno("danilo"))
	time.Sleep(time.Second * 5)
	fmt.Println(retorno("gabriel"))

	// ejemplo 2
	canal := make(chan string)
	go func() {
		canal <- retorno("hola")
	}()
	fmt.Println(<-canal)
}

func retorno(parametro string) string {
	return "hola  " + parametro
}
