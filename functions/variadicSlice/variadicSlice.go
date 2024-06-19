package main

import "fmt"

func printApproved(approveds ...string) {
	fmt.Println("Approved Students")
	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	approveds := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	printApproved(approveds...)
}
