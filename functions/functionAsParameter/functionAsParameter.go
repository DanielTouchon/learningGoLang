package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func exec(function func(int, int) int, n1, n2 int) int {
	return function(n1, n2)
}

func main() {
	result := exec(multiply, 3, 4)
	fmt.Println(result)
}
