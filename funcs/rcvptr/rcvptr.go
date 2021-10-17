package main

import (
	"fmt"
)

/* Go is a 100% Call by value language */

func main() {

	fmt.Println("------------Passing int pointer to a function---------")
	i := 2
	x := &i
	fmt.Println("retptr: &i:", x)
	fmt.Println("retptr: i:", *x)
	fmt.Println("retptr: &x:", &x)
	fmt.Println("retptr: result:", retptr(x))
	fmt.Println("-------------------------------------------------------")

	fmt.Println("------------Passing map to a function---------")
	m := make(map[int]string)
	addMap(m)
	fmt.Printf("main: &m:%p\n", &m)
	fmt.Println("main: m:", m)
	fmt.Println("-------------------------------------------------------")

	fmt.Println("------------Passing slice to a function---------")
	s := make([]int, 0)
	appSlc(s)
	fmt.Printf("main: &s:%p; len(s):%d; cap(s):%d\n", &s, len(s), cap(s))
	Show("main:s:", s)
	fmt.Println("-------------------------------------------------------")
}

func appSlc(s []int) {
	s = append(s, 1, 2, 3)
	// for i, _ := range s {
	// 	s[i] = i
	// }
	fmt.Printf("appSlc: &s:%p; len(s):%d; cap(s):%d\n", &s, len(s), cap(s))
	Show("appSlc: s:", s)
}

func appSlcRef(s *[]int) {
	// []int{0,0,0,0}
	*s = append(*s, 1, 2, 3)

	for i, _ := range *s {
		(*s)[i] = i
	}

	fmt.Printf("appSlc: &ab:%p; len(ab):%d; cap(ab):%d\n", &ab, len(ab), cap(ab))
	Show("appSlc: s:", s)
	// Show("appSlc: ab:", ab)
}

func addMap(m map[int]string) {
	m[1] = "first"
	fmt.Printf("addMap: &m:%p\n", &m)
	fmt.Println("addMap: m:", m)
}

func retptr(x *int) int {
	y := *x * *x
	fmt.Println("retptr: &i:", x)
	fmt.Println("retptr: i:", *x)
	fmt.Println("retptr: &x:", &x)
	fmt.Println("retptr: result:", y)
	return y
}
