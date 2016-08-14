package main

import (
	"fmt"
)

// I reused this code from github.com/rlmcpherson
// It was a package created for testing a Clock object
// I changed the package name and added a main function

// Create a data structure of type Clock with one field
type Clock struct {
	min int
}

// Time is a method of the Clock struct
// It accepts two ints and returns a Clock
// It converts hour into minutes then adds minute
// If Clock.min is negative, add 24 hours
// This method is not required for my basis clock
// I left it here as an exercise - see below
func Time(hour, minute int) Clock {
	clock1 := Clock{(60*hour + minute) % 1440}
	if clock1.min < 0 {
		clock1.min += 1440
	}
	return clock1
}

// Similar to above, accepts two ints and returns a Clock
// Outputs the number of minutes added
func Add(time, min int) Clock {
	c := Clock{time}
	c.min = (c.min + min) % 1440
	fmt.Println("Minutes added", min)
	if c.min < 0 {
		c.min += 1440
	}
	return c
}

// String() is a stringer, not sure how they work
// golang.org says it uses a command called generate
// It automatically formats calls to print
func (clock1 Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock1.min/60, clock1.min%60)
}

func main() {
	clock1 := Clock{1400}
	clockTime := clock1.min
	// fmt.Println("Clock time:", clockTime)
	fmt.Println("Clock 1 is:", clock1)
	// myTime := Time(16, -15)
	// fmt.Println("My time is:", myTime)
	changeClock := Add(clockTime, -30)
	fmt.Println("New time is", changeClock)
}
