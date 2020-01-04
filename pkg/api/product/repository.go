package product

type Repository interface {
	GetProductList(acceptLang string) []Product
}

type repository struct {
	Ds DataSource
}

func (r *repository) GetProductList(acceptLang string) []Product {
	return r.Ds.GetList(acceptLang)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
