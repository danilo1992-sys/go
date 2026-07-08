package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("La hora actual es: ", time.Now())

	fecha := time.Now()

	fmt.Println("The year is:", fecha.Year())
}
