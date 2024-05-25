package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // getting the memory adress of the variable

	*p++ // incrementing the variable to which the pointer points
	i++

	fmt.Println(p, *p, i, &i)

	//p++
	// there is no pointer arithmetics in go
}
