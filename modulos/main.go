package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("La hora actual es: ", time.Now())

	fecha := time.Now()

	fmt.Println("The year is:", fecha.Year())
	fmt.Println("The month is:", fecha.Month())
	fmt.Println("The month is:", int(fecha.Month()))
	fmt.Println("The day is:", fecha.Day())
	fmt.Println("The hour is:", fecha.Hour())
	fmt.Println("The minutes is:", fecha.Minute())
	fmt.Println("The second is:", fecha.Second())
	fmt.Printf("%v/%v/%v \n", fecha.Day(), int(fecha.Month()), fecha.Year())
	fmt.Printf("%v:%v:%v \n", fecha.Hour(), int(fecha.Minute()), fecha.Second())
	ahora := time.Now()
	fmt.Println("Mas 22 dias: ")
	fecha1 := ahora.Add(time.Hour * 24 * 22)
	fmt.Printf("%v/%v/%v \n", fecha1.Day(), int(fecha1.Month()), fecha1.Year())
	fmt.Println("Restar")
	fmt.Println("Restar 22 dias: ")
	fecha2 := ahora.Add((time.Hour * 24 * 22) * -22)
	fmt.Printf("%v/%v/%v \n", fecha2.Day(), int(fecha2.Month()), fecha2.Year())
	fmt.Println("Dentro de 1 año: ")
	fecha3 := ahora.Add(365 * 24 * time.Hour)
	fmt.Printf("%v/%v/%v \n", fecha3.Day(), int(fecha3.Month()), fecha2.Year())
	fmt.Println("")
	fmt.Println(FormatoFecha(ahora))
}

func FormatoFecha(fecha time.Time) string {
	v := fmt.Sprintf("%v/%v/%v \n", fecha.Day(), int(fecha.Month()), fecha.Year())
	return v
}
