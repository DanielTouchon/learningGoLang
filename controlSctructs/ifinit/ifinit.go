package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := randNumber(); i >= 5 {
		fmt.Println(i)
		fmt.Println("Win!")
	} else {
		fmt.Println(i)
		fmt.Println("Lost!")
	}
}
