package main

import (
	"log"
	"os"
)

func main() {
	//log.Fatal("Error fatal en el servidor")
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf("Error linea %v", 1)
}
