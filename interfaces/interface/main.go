package main

import "fmt"

type toy struct {
	name  string
	price float64
}

func (t toy) print() {
	fmt.Printf("Toy: %+v\n", t)
}

func main() {
	var (
		b = book{name: "Harry potter", price: 250.50}
		g = game{name: "Dungeons & dragons", price: 150.50}
		p = puzzle{name: "Crossword", price: 50.50}
		t = toy{name: "barbie", price: 350.50}
	)

	var store items
	store = append(store, b, g, p, t)
	store.print()
}
