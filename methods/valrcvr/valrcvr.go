package main

import (
	"fmt"
	"unsafe"
)

/*
	Sample output of this program:
	$ go run valrcvr.go
	0xc000120000
	size of h:80000
*/
type huge struct {
	Items [10000]int
}

func (h huge) size() {
	fmt.Printf("%p\n", &h)
	fmt.Printf("size of h:%v\n", unsafe.Sizeof(h))
}

func main() {
	var h huge
	h.size()
}
