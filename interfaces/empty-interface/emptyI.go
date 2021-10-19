package main

import "fmt"

func main() {

	fmt.Println("Result:", Add(1.0, 2.1))
	fmt.Println("Result:", Add("sahil", "wale"))
}

func Add(a, b interface{}) interface{} {
	switch val := a.(type) {
	case string:
		return val + "-" + b.(string)
	case float64:
		return val + b.(float64)
	}
	return "no case matched!!"
}
