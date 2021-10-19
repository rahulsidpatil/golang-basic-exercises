package main

import "fmt"

type puzzle struct {
	name  string
	price float64
}

func (p puzzle) print() {
	fmt.Printf("Puzzle: %+v\n", p.name)
}
