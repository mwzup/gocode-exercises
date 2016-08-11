package main

import (
	"fmt"
)

// Clock exercise from http://exercism.io/exercises/go/clock/readme
// Implement a clock that handles times without dates.
// Create a clock that is independent of date.
// You should be able to add and subtract minutes to it.

// addTime accepts two ints and  returns the sum
func addTime(time, min int) int {
	clock1 := time + min
	fmt.Println("Minutes added", min)
	if clock1 < 0 {
		clock1 += 1440
	}
	return clock1
}

// clockString converts time int to formatted string
func clockString(time int) string {
	return fmt.Sprintf("%02d:%02d", time/60, time%60)
}

func main() {
	// clockTime sets the initial value for Clock in minutes
	// clockTime can be set from user input
	clockTime := 800
	fmtTime := clockString(clockTime)
	fmt.Println("Clock time:", fmtTime)
	// changeClock sets the minutes to be added
	// changeClock can be set from user input
	changeClock := addTime(clockTime, 146)
	newTime := clockString(changeClock)
	fmt.Println("New time is", newTime)
}
