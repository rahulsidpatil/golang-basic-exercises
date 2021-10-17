package main

import (
	"errors"
	"fmt"
)

type list []int

func (l *list) print() {
	fmt.Printf("list:%#v\n", l)
}

func (l *list) Add(n int) *list {
	*l = append((*l), n)
	return l
}

// list = {1,2,3}
// list1 = {4,5,6}
func (l *list) AddAll(l1 list) *list {
	*l = append((*l), l1...)
	return l
}

// list = {1,2,3,4,5,6,7} => remove(3) => {1,2,4,5,6,7}
// Equivalence partitioning
// equivalance class 1: boundary values=> {1,7} => a user may want to remove boundary elements
// equivalance class 2: non-boundary values => {2,3,4,5,6,} => a user may want to remove non-boundary elements
// equivalance class 3: absent elements => {0, 10, 1000.....}
func (l *list) Remove(n int) (*list, error) {
	var present bool
	// values present in the list (boundary/non-boundary values)
	for i, e := range *l {
		if e == n {
			temp := (*l)[0:i]
			result := append(temp, (*l)[i+1:]...)
			return &result, nil
		}
	}
	//absent elements
	if !present {
		return l, errors.New("you are trying to remove a non-existing element")
	}
	return l, nil
}
