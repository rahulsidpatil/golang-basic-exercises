package main

import "fmt"

func main() {
	fmt.Println("Sum:", sumIt(1, 2, 3))

	// slc1 := []int{1, 2, 3, 4, 5}
	// slc := make([]int, 0)
	// slc = append(slc, slc1...) // unpack operator
	// fmt.Println("Slc:", slc)
}

func sumIt(input ...int) int { // pack operator (syntax: ...<data type>)
	var sum int
	for _, i := range input {
		sum += i
	}
	return sum
}
