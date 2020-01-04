package api

import (
	"github.com/prongbang/gomtl/pkg/api/product"
	"github.com/prongbang/gomtl/pkg/database"
)

func CreateAPI(dbSource database.DataSource) API {
	route := product.NewProductAPI(dbSource)
	apiAPI := NewAPI(route)
	return apiAPI
}
