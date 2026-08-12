package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	random := rand.Intn(101)
	fmt.Println(random)
	min := 10
	max := 100
	rand.Seed(time.Now().UnixNano())
	random2 := rand.Intn(max-min) + min
	fmt.Println(random2)
}
