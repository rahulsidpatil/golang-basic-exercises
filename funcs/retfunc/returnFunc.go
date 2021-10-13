package main

import "fmt"

func main() {
	i := returnFunc()
	j := returnFunc()

	fmt.Println("i1", i())
	fmt.Println("i2", i())
	fmt.Println("i3", i())
	fmt.Println("i4", i())
	fmt.Println("j1", j())
	fmt.Println("j2", j())
	fmt.Println("j3", j())
	fmt.Println("j4", j())
}

func returnFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
