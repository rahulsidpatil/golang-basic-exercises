package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Time Multiplier
//
//  You should get an argument from the command-line and
//  you need to multiply the time duration value `t` with
//  the given argument.
//
//  1. Get an argument from the command-line
//  2. Convert it to int64 and store it in a variable
//  3. Multiply the `t` variable with that variable
//  4. Print the result
//
// HINT
//  You can use ParseInt to convert the command-line
//    argument into an int64 value.
//
//  You can skip the error value using a blank-identifier.
//
// EXPECTED OUTPUT
//
//  When runned like this:
//    go run main.go 2
//
//  It should print this:
//    3h0m0s
// ---------------------------------------------------------

func main() {
	// DONT TOUCH THIS
	// 1h30m means: 1 hour 30 minutes
	t, _ := time.ParseDuration("1h30m")

	// TYPE YOUR CODE HERE
	// ....

	// DONT TOUCH THIS
	fmt.Println(t)
}
package main

// ---------------------------------------------------------
// EXERCISE: Try the scopes
//
//  1. Create two files: main.go and printer.go
//
//  2. In printer.go:
//     1. Create a function named hello
//     2. Inside the hello function, print your name
//
//  3. In main.go:
//     1. Create the usual func main
//     2. Call your function just by using its name: hello
//     3. Create a function named bye
//     4. Inside the bye function, print "bye bye"
//
//  4. In printer.go:
//     1. Call the bye function from
//        inside the hello function
// ---------------------------------------------------------

func main() {
}