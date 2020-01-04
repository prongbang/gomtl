package api

import (
	"github.com/prongbang/gomtl/pkg/api/product"
	"net/http"
)

type API interface {
	Register()
}

type api struct {
	ProductRoute product.Route
}

func (a *api) Register() {
	a.ProductRoute.Initial()

	_ = http.ListenAndServe(":8080", nil)
}

func NewAPI(
	productRoute product.Route,
) API {
	return &api{
		ProductRoute: productRoute,
	}
}
