package main

import (
	"errors"
	"fmt"
)

// stack - LIFO (Last in First Out)
type stack []int

// Test Driven Development(TDD)
//1) Create a method signature
//2) Create a unit test for the method
//3) Run the unit test. It will fail
//4) Add logic to the method in step 1
//5) Run the unit test in step 2). It should pass

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() (poped int, err error) {
	if len(*s) == 0 {
		err = errors.New("trying to pop from an empty stack")
		return
	}

	poped = (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	fmt.Printf("Error:%v\n", err)
	return
}

func (s *stack) printStack() {
	fmt.Println("Present stack:", s)
}

// s=> s{}
// s.push(3)->push(4)->push(6)->push(9)-> => s{3,4,6,9}
//s.pop()=> s{3,4,6}

func main() {
	var s stack
	s.printStack()
	// err := s.pop()
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }
}
