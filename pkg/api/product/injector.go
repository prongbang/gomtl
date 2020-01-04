package product

import "github.com/prongbang/gomtl/pkg/database"

func NewProductAPI(dbSource database.DataSource) Route {
	productDataSource := NewDataSource(dbSource)
	productRepository := NewRepository(productDataSource)
	productUseCase := NewUseCase(productRepository)
	productHandler := NewHandler(productUseCase)
	productRoute := NewRoute(productHandler)
	return productRoute
}
