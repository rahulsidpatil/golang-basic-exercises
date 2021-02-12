package library

//file scope
import "fmt"

//package scope
var gopher = "gopher!!"

func bye() {
	//block scope
	var msg = "bye bye"
	fmt.Println("from library hello.go:", msg, gopher)
}
