package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Dwight Schrute": 1325.45,
		"James Halpert":  3542.51,
		"Angela Lipton":  5512.34,
	}

	funcsESalarios["Kevin Malone"] = 2535.14
	delete(funcsESalarios, "Non-existent") //doesn't crash the code, even if element to be deleted doesn't exist

	for name, wage := range funcsESalarios {
		fmt.Println(name, wage)
	}

	for _, wage := range funcsESalarios {
		fmt.Println(wage)
	}

	for wage := range funcsESalarios {
		fmt.Println(wage) //if i want to only one variable in this format, it necessary means I will use the keys, not the values
	}
}
