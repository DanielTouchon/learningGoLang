package main

import "fmt"

func main() {
	approved := make(map[int]string) //maps must always be initialized
	approved[1] = "John"
	approved[2] = "Mary"
	approved[3] = "Jane"
	fmt.Println(approved)

	for id, name := range approved {
		fmt.Printf("%s (ID: %d)\n", name, id)
	}

	for _, name := range approved {
		fmt.Printf("%s\n", name)
	}

	fmt.Println(approved[2])

	delete(approved, 1)

	fmt.Println("Teste", approved[1])
}
