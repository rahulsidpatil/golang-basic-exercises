package main

import "fmt"

func main() {
	var l1 list
	l1.Add(4).Add(5).Add(6)
	l1.print()
	var l list
	l.Add(1).Add(2).Add(3)
	l.print()
	l.AddAll(l1)
	l.print()

	res, err := l.Remove(3)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
	}
	res.print()
}
