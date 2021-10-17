package main

import (
	"fmt"
	"unsafe"
)

type huge struct {
	Items [10000]int
}

/*
	Sample output of this program:
	$ go run ptrrcvr.go
	0xc00000e028
	size of h:8
*/
func (h *huge) size() {
	fmt.Printf("%p\n", &h)
	fmt.Printf("size of h:%v\n", unsafe.Sizeof(h))
}

func main() {
	var h huge
	h.size()
}
