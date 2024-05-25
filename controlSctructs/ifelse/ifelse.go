package main

import "fmt"

func getResult(score float64) {
	if score >= 7 {
		fmt.Println("Approved!\nScore:", score)
	} else if score >= 4 && score < 7 {
		fmt.Println("Summer School!\nScore:", score)
	} else {
		fmt.Println("Reproved!\nScore:", score)
	}
}

func main() {
	getResult(7.3)
	getResult(5.1)
	getResult(3.9)
}
