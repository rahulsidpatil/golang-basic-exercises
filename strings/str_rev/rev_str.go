package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	str := "a1b2c3d"
	fmt.Println("Original string:", str)

	fmt.Println("Rverse the string:", RevChar(str))

	fmt.Println("Rverse only letters in the string:", RevStr(str))
}

// RevChar ... Reverse only the letters in the given string
// Example:
// input: "a1b2c3d"
// output:"d1c2b3a"
func RevChar(str string) string {
	strlen := utf8.RuneCountInString(str)
	r := []rune(str)
	for i, s := range str {
		if _, err := strconv.ParseInt(string(s), 10, 64); err != nil {
			if _, err := strconv.ParseInt(string(str[strlen-i-1]), 10, 64); err != nil {
				r[strlen-i-1] = s
				r[i] = rune(str[strlen-i-1])
			} else {
				r[i] = s
			}
		} else {
			r[i] = s
		}
	}
	return string(r)
}

// RevStr ... Reverse the given string
// Example:
// input: "a1b2c3d"
// output:"d3c2b1a"
func RevStr(str string) (result string) {
	// var result string
	for _, s := range str {
		result = string(s) + result
	}
	return result
}
