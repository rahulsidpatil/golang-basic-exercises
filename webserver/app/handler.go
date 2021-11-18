package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Name string
	ID   int
}

type Store struct {
	Products []Product
}

func GetStore() *Store {
	return &Store{}
}
func (s *Store) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if s.Products == nil {
			rw.Write([]byte(StoreEmptyMsg))
			return
		} else {
			products, _ := json.Marshal(s.Products)
			rw.Write(products)
			return
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var newProduct Product
		err = json.Unmarshal(body, &newProduct)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if newProduct.ID == 0 {
			rw.WriteHeader(http.StatusBadRequest)
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
		return
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var newProduct Product
		err = json.Unmarshal(body, &newProduct)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if newProduct.ID == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		for i, p := range s.Products {
			if newProduct.ID == p.ID {
				s.Products[i].Name = newProduct.Name
				rw.WriteHeader(http.StatusOK)
				return
			}
		}
	case http.MethodDelete:
		urlPathParams := strings.Split(r.URL.Path, "products/")
		productID, err := strconv.Atoi(urlPathParams[len(urlPathParams)-1])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		for i, p := range s.Products {
			if productID == p.ID {
				temp := s.Products[:i]
				s.Products = append(temp, s.Products[i+1:]...)
				rw.WriteHeader(http.StatusOK)
				return
			}
		}
		rw.WriteHeader(http.StatusNotModified)
		return
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
