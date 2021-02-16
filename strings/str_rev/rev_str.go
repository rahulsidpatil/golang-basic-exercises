package main

import "fmt"

func main() {
	str := "abcde"
	fmt.Println("Before:", str)

	fmt.Println("After:", revstr(str))
}

func revstr(str string) (result string) {
	for _, s := range str {
		result = string(s) + result
	}
	return result
}
