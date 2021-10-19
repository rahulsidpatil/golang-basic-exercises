package main

import "fmt"

type book struct {
	name  string
	price float64
}

func (b book) print() {
	fmt.Printf("Book: %+v\n", b)
}
