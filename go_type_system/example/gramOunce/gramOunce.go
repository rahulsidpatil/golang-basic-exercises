package main

import (
	"fmt"
	"strconv"
)

// gram => defined/named type; float64 => source type; float64 => underlying type
type gram float64

// ounce => defined/named type; float64 => source type; float64 => underlying type
type ounce float64

type gm = float64
type on = float64

const mfactor = 0.035274

func main() {
	s := make([]interface{}, 2)
	s[0], s[1] = strconv.Atoi("43")
	fmt.Printf("\n%v", s)
	// var g gram = 1000
	// var o ounce
	// o = ounce(g) * mfactor // a Defined type inherits underlying type's operations
	// fmt.Printf("%g grams is: %.2f ounces\n", g, o)

	// // How it works with the Aliased types
	// var gr gm = 1000
	// var ou on
	// ou = gr * mfactor // an Aliased type is same as it's source type
	// fmt.Printf("%g gm is: %.2f ou\n", gr, ou)
}
