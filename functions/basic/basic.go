package main

import "fmt"

func f1() {
	fmt.Println("F1: R1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "R3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("R4: %s %s\n", p1, p2)
}

func f5() (string, string) {
	return "Return 1", "Return 2"
}

func main() {
	f1()
	f2("Param 1", "Param 2")
	r3, r4 := f3(), f4("Param 1", "Param 2")
	fmt.Printf("F3: %s\nF4: %s", r3, r4)
	//r51,_ := f5()
	//fmt.Printf("F5: %s",r52)
	r51, r52 := f5()
	fmt.Printf("F5: %s, %s", r51, r52)
}
