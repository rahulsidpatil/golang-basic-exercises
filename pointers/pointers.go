package main

import "fmt"

func main() {
	var c int = 24
	p := &c
	v := *p

	fmt.Printf("Type of c:%T\n", c)
	fmt.Printf("Value of c:%v\n", c)

	c++
	fmt.Printf("Type of p:%T\n", p)
	fmt.Printf("Value at address p:%v\n", *p)

	fmt.Printf("Type of v:%T\n", v)
	fmt.Printf("Value of v:%v\n", v)
}
