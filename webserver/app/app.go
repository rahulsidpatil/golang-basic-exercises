package app

import (
	"log"
	"net/http"
)

func InitServer(s *Store) {
	http.HandleFunc("/store", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(WelcomeMsg))
	})
	http.Handle("/store/products", s)
	http.Handle("/store/products/", s)
	log.Fatal(http.ListenAndServe(":5001", nil))
}
