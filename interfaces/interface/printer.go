package main

import "fmt"

type printer interface {
	print()
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
