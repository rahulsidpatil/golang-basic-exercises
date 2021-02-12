package main

//file scope
import (
	"fmt"

	"github.com/rahulsidpatil/golang-basic-exercises/packages_n_scopes/example/library"
)

//package scope
var gopher = "gopher!!"

func main() {
	hello()
	library.Hello()
}
func bye() {
	//block scope
	var msg = "bye bye"
	fmt.Println("from main.go:", msg, gopher)
}
