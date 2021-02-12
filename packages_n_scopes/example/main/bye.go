package main

//file scope
import "fmt"

func hello() {
	//block scope
	var msg = "hello "
	fmt.Println("from bye.go:", msg, gopher)
	bye()
}
