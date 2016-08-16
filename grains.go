package main

import (
    "fmt"
    "math"
) 

// Grains
// Write a program that calculates the number of grains of wheat 
// on a chessboard given that the number on each square doubles.
// There are 64 squares on a chessboard.
// Write a program that shows:
//  a. how many grains were on each square
//  b. the total number of grains

func main() {
   
    // squares is a slice used to hold the value of each 
    // 'square' on the chessboard.
    // sum is used to calculates the values.
	var sum uint64 = 1
	squares := []uint64{1}
	for i := 0; i < 64; i++ {
		sum *= 2
        squares = append(squares, sum)
        fmt.Println(i+1 ,squares[i])    
    }

    // total calculates the total number 
    // of wheat grains on the chessboard. 
    total := uint64(math.Pow(2, float64(65))) - 1
    fmt.Printf("Total grains of wheat: %v",total)   
}