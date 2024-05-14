package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(97))

	fmt.Println("Teste " + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	b, _ = strconv.ParseBool("false")
	if !b {
		fmt.Println("Falso")
	}

	b, _ = strconv.ParseBool("0")
	if !b {
		fmt.Println("Falso")
	}

	b, _ = strconv.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}

	b, _ = strconv.ParseBool("2")
	if !b {
		fmt.Println("Falso")
	}
}
