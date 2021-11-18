package main

import "github.com/rahulsidpatil/golang-basic-exercises/webserver/app"

func main() {
	app.InitServer(app.GetStore())
}
