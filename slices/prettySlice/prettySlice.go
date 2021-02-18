package main

import (
	sl "github.com/rahulsidpatil/prettyslice"
)

func main() {

	s := make([]string, 3)
	sl.Show("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	sl.Show("set:", s)
	sl.Show("get:", s[2])

	sl.Show("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	sl.Show("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	sl.Show("cpy:", c)

	l := s[2:5]
	sl.Show("sl1:", l)

	l = s[:5]
	sl.Show("sl2:", l)

	l = s[2:]
	sl.Show("sl3:", l)

	t := []string{"g", "h", "i"}
	sl.Show("dcl:", t)
}
