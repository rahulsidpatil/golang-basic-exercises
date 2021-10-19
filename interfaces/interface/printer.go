package main

import "fmt"

type printer interface {
	print()
}

type discounter interface {
	discount(float64)
}

type items []printer

func (i items) print() {
	if len(i) == 0 {
		fmt.Printf("Sorry!! store empty")
		return
	}
	for _, it := range i {
		it.print()
	}
}

func (i items) discount() {
	if len(i) == 0 {
		fmt.Printf("Sorry!! store empty")
		return
	}
	for _, it := range i {
		if g, ok := it.(discounter); ok {
			g.discount(5)
		}
	}
}
