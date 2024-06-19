package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Agerage: %2f", average(7.7, 8.1, 5.9, 6.0))
}
