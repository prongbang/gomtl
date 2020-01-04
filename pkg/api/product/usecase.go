package product

type UseCase interface {
	GetProductList(acceptLang string) []Product
}

type useCase struct {
	Repo Repository
}

func (u *useCase) GetProductList(acceptLang string) []Product {
	return u.Repo.GetProductList(acceptLang)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
