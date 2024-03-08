package handlers

import (
	"GoWeb/proyect/internal"
	"github.com/bootcamp-go/web/response"
	"net/http"
)

func NewProductDefault(service internal.ProductService) *ProductDefault {
	return &ProductDefault{service: service}
}

type ProductDefault struct {
	service internal.ProductService
}

type ProductHandlerGetAll struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type responseGetAllProducts struct {
	Message string                  `json:"message"`
	Data    []*ProductHandlerGetAll `json:"data"`
	Error   bool                    `json:"error"`
}

func (p *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := p.service.GetAll()

		code := http.StatusOK
		data := make([]*ProductHandlerGetAll, 0)
		body := &responseGetAllProducts{Message: "OK", Data: data, Error: false}
		for _, product := range products {
			body.Data = append(body.Data, &ProductHandlerGetAll{
				Id:          product.Id,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration,
				Price:       product.Price,
			})
		}
		response.JSON(w, code, body)
	}
}

func (p *ProductDefault) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong!"))
	}
}

//func (p *ProductDefault) Get() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(p.service)
//	}
//}

//func (p *ProductDefault) GetByID() http.HandlerFunc {
//return func(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
//	for _, product := range h.data {
//		if product.Id == id {
//			json.NewEncoder(w).Encode(product)
//			return
//		}
//	}
//	w.WriteHeader(http.StatusNotFound)
//	json.NewEncoder(w).Encode(MyResponse{Message: "Product not found"})
//}
//}

//func (p *ProductDefault) Search() http.HandlerFunc {
//return func(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
//
//	newSlice := []Product{}
//	for _, product := range h.data {
//		if product.Price < price {
//			newSlice = append(newSlice, product)
//		}
//	}
//
//	if h.data == nil {
//		w.WriteHeader(http.StatusNotFound)
//		json.NewEncoder(w).Encode(MyResponse{Message: "Product not found"})
//		return
//	}
//
//	json.NewEncoder(w).Encode(newSlice)
//	return
//
//}
//}
