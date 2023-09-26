package interfaces

import "net/http"

type ProductsHandlers interface {
	GetHomeProducts(w http.ResponseWriter, r *http.Request)
	PostPublishProduct(w http.ResponseWriter, r *http.Request)
	GetPublishProductTemplate(w http.ResponseWriter, r *http.Request)
}
