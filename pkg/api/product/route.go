package product


import (
	"net/http"
)

type Route interface {
	Initial()
}

type route struct {
	Handle Handler
}

func (r *route) Initial() {
	http.HandleFunc("/v1/product", r.Handle.GetProduct)
}

func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}