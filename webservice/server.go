package main

import (
	"log"
	"net/http"
)

type helloHandler struct {
	Msg string
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Msg))
}

func middlewarefunc(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//actual middleware code goes here
		w.Write([]byte("middleware called!!\n"))
		handler.ServeHTTP(w, r)
	})
}

func main() {
	h := &helloHandler{Msg: "Hello client!!"}
	http.HandleFunc("/svc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("webservice called using handleFUnc!!"))
	})
	http.Handle("/hello", middlewarefunc(h))
	log.Fatal(http.ListenAndServe(":5000", nil))
}
