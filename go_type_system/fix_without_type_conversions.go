package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: fix without type conversion
//
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
//
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------

func main() {
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)
}
