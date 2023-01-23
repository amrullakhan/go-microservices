package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/amrullakhan/go-microservices/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		p.updateProduct(w, r)
		return
	}

	if r.Method == http.MethodDelete {
		p.deleteProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to process the payload", http.StatusBadRequest)
	}

	data.AddProduct(product)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product Created Successfull"))
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	matcher := regexp.MustCompile(`/([0-9]+)`)
	grabbed := matcher.FindAllStringSubmatch(r.URL.Path, -1)

	if len(grabbed) != 1 {
		http.Error(w, "Inalid URI", http.StatusBadRequest)
		return
	}

	if len(g[0]) != 2 {
		http.Error(w, "Inalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Inalid URI", http.StatusBadRequest)
		return
	}

	product := &data.Product{}
	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to process the payload", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("Updated the product"))
}

func (p *Products) deleteProduct(w http.ResponseWriter, r *http.Request) {

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
