package main

import "fmt"

func main() {
	//Map initialization
	m := make(map[string]int, 2)
	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 15
	m["k4"] = 17
	// fmt.Printf("\nmap:%+v; what is map:%T; \n", m, m)
	fmt.Println("m:", m)

	delete(m, "k1")
	fmt.Printf("\nmap:%+v\n\n", m)
	delete(m, "k1")

	// m1 := map[string]int{
	// 	"p1": 1,
	// 	"p2": 2,
	// }
	// fmt.Println("m1:", m1)

	// m1 = m

	// fmt.Println("m1(afterr):", m1)

	// for key, val := range m {
	// 	fmt.Println(key, val)
	// }

	// // v1 := m["k1"]
	// // fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	// // delete(m, "k2")
	// // fmt.Println("map:", m)

	// // _, prs := m["k2"]
	// // fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
