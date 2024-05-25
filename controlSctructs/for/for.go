package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Forever printing...")
		time.Sleep(time.Second * 5) //this is how you get an inifinte loop in go
	}
}
