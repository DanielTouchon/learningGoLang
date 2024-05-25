package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Bom dia")
	case t.Hour() < 24:
		fmt.Println("Bom dia")
	}
}
