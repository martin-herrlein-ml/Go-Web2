package main

import (
	"GoWeb/proyect/internal/handlers"
	"GoWeb/proyect/internal/repository"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//type MyHandler struct {
//	data []Product
//}

//func NewHandler() *MyHandler {
//	products, err := readJsonFile()
//	if err != nil {
//		println(err.Error())
//	}
//	return &MyHandler{data: products}
//}

func main() {

	stBase := repository.NewStorageJsonFile("products.json")
	st := repository.NewProductSlice(stBase)

	ct:= handlers.NewProductDefault(st)

	productsJson := repository.NewStorageJsonFile("products.json")
	productSlice :=
	rp := repository.NewProductSlice(productsJson)
	r := chi.NewRouter()

	h := NewHandler()

	r.Get("/ping", h.ping())
	r.Get("/products", h.GetAll())
	r.Get("/products/{id}", h.GetByID())
	//example endpoint http://localhost:8080/products/search?price=100
	r.Get("/products/search", h.Search())

	http.ListenAndServe(":8080", r)

}

type MyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
