package library

//file scope
import "fmt"

// Hello ... say hello
func Hello() { // exported function - public scope
	//block scope
	var msg = "hello "
	fmt.Println("from library bye.go:", msg, gopher)
	bye()
}
