package main

import (
	"fmt"
)

// Raindrops
// Write a program that converts a number to a string,
// the contents of which depends on the number's factors.
// If the number contains 3 as a factor, output 'Pling'.
// If the number contains 5 as a factor, output 'Plang'.
// If the number contains 7 as a factor, output 'Plong'.
// If the number does not contain 3, 5, or 7 as a factor,
// just pass the number's digits straight through.

// Test samples
// 34  (2*17)  = 34
// 9   (3*3)   = Pling
// 15  (3*5)   = PlingPlang
// 105 (3*5*7) = PlingPlangPlong

func main() {

	var rainDrops string = ""
	var num int

	fmt.Printf("Enter a number: ")
	fmt.Scan(&num)

	// if number is divisible by 3 or 5 or 7, add text to string
	// 3 = Pling : 5 = Plang : 7 = Plong

	if num%3 == 0 {
		rainDrops = "Pling"
	}
	if num%5 == 0 {
		rainDrops = rainDrops + "Plang"
	}
	if num%7 == 0 {
		rainDrops += "Plong"
	}

	if len(rainDrops) > 0 {
		fmt.Println(rainDrops)
	} else {
		fmt.Println(num)
	}

	//// this works
	// if rainDrops != "" {
	// 	fmt.Println(rainDrops)
	// } else {
	// 	fmt.Println(num)
	// }

	//// this works
	// if rainDrops == "" {
	// 	fmt.Println(num)
	// } else {
	// 	fmt.Println(rainDrops)
	// }

}
