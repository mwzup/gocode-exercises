package main

import (
	"fmt"
)

// isLegit compares triangle side lengths for triangle inequality
// and returns true or false
func isLegit(a, b, c float64) bool {
	legit := false
	if a+b >= c && b+c >= a && a+c >= b {
		legit = true
	} else {
		legit = false
	}
	return legit
}

// triangleType determines which type of
// triangle is formed from the three side lengths
func triangleType(a, b, c float64) string {
	result := ""
	switch {
	// All sides are equal
	case a == b && a == c:
		result = "The triangle is equilateral"
	// Two sides are equal
	case a == b || a == c || b == c:
		result = "The triangle is isosceles"
	// No sides are equal
	default:
		result = "The triangle is scalene"

	}
	return result
}

func main() {
	var a, b, c float64
	var result string

	// request and store user input for triangle side lengths
	fmt.Println("\nDetermine the type of a triangle")
	fmt.Printf("Enter length of side A: ")
	fmt.Scan(&a)
	fmt.Printf("Enter length of side B: ")
	fmt.Scan(&b)
	fmt.Printf("Enter length of side C: ")
	fmt.Scan(&c)

	legit := isLegit(a, b, c)

	if legit {
		result = triangleType(a, b, c)
	} else {
		result = "Triangle can not exist."
	}

	fmt.Println(result)
}
