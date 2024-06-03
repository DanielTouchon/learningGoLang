package main

import "fmt"

func main() {
	employeeByInitialLetter := map[string]map[string]float64{
		"G": {
			"Gary Neville": 15456.78,
			"Gerd Muller":  23745.12,
			"George Best":  9371.09,
			"Gareth Bale":  51970.00,
		},
		"L": {
			"Lionel Messi":    107281.90,
			"Luka Modric":     97987.16,
			"Leighton Baines": 16029.55,
			"Leroy Sané":      28090.38,
		},
		"R": {
			"Robert Lewandowski": 76378.09,
			"Rodri":              20987.12,
			"Ruben Dias":         87371.00,
			"Ronald Araujo":      15609.21,
		},
	}

	delete(employeeByInitialLetter, "R")

	for letter, employees := range employeeByInitialLetter {
		fmt.Printf("Employees starting with '%s'\n\n", letter)
		for employee, wage := range employees {
			fmt.Println(employee, wage)
		}
	}
}
