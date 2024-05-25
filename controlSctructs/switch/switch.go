package main

import "fmt"

func scoreConcept(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7, 6, 5, 4, 3, 2, 1, 0:
		return "B"
	default:
		return "Invalid score"
	}
}

func main() {
	fmt.Println(scoreConcept(9.8))
	fmt.Println(scoreConcept(6.9))
	fmt.Println(scoreConcept(2.1))
}
