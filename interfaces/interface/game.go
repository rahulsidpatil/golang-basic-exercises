package main

import "fmt"

type game struct {
	name  string
	price float64
}

func (g game) print() {
	fmt.Printf("Game: %+v\n", g)
}

func (g *game) discount(percent float64) {
	g.price *= (percent * 0.01)
}
