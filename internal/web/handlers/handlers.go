package handlers

import "net/http"

type ProductsHandlers interface {
	GetAndSearchProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	PostProduct(w http.ResponseWriter, r *http.Request)
	DeleteProductById(w http.ResponseWriter, r *http.Request)
	PutProduct(w http.ResponseWriter, r *http.Request)
}
