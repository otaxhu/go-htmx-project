package std_http_interfaces

import "net/http"

type Handlers interface {
	HomeGetProducts(w http.ResponseWriter, r *http.Request)
}
