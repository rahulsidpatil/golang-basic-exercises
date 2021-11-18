package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	WELCOME_MSG     = "Welcome to the online store!!!"
	STORE_EMPTY_MSG = "The store is empty. try later!!"
)

type Product struct {
	Name string
	ID   int
}

type Store struct {
	Products []Product
}

func (s *Store) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if len(s.Products) == 0 {
			rw.Write([]byte(STORE_EMPTY_MSG))
			return
		} else {
			data, err := json.Marshal(s)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
				return
			}
			rw.Write([]byte(data))
		}
	case http.MethodPost:
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		newProduct := Product{}
		err = json.Unmarshal(data, &newProduct)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		for _, p := range s.Products {
			if newProduct.ID == p.ID {
				rw.WriteHeader(http.StatusConflict)
				return
			}
		}
		s.Products = append(s.Products, newProduct)
		rw.WriteHeader(http.StatusCreated)

	case http.MethodDelete:
		pathParams := strings.Split(r.URL.Path, "products/")
		pid, err := strconv.Atoi(pathParams[len(pathParams)-1])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		for i, p := range s.Products {
			if pid == p.ID {
				items := s.Products[:i]
				s.Products = append(items, s.Products[i+1:]...)
				rw.WriteHeader(http.StatusOK)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(WELCOME_MSG))
	})
	s := &Store{Products: []Product{{"soap", 1}}}
	http.Handle("/store/products", s)
	http.Handle("/store/products/", s)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
