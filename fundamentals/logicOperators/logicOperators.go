package main

import "fmt"

func shopping(job1, job2 bool) (bool, bool, bool) {
	buy50InchTv := job1 && job2
	buy32InchTv := job1 != job2 // there's no XOR in GO, this solution emulates it (limited to scenarios when you can guarantee both variables hold bool)
	buyIceCream := job1 || job2

	return buy50InchTv, buy32InchTv, buyIceCream
}

func main() {
	tv50, tv32, iceCream := shopping(true, true)
	fmt.Printf("50-inch:%t\n32-inch:%t\nIce Cream:%t\nHealthy:%t", tv50, tv32, iceCream, !iceCream)
}
