package usecase

import (
	"products/model"
	"products/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProdutUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProducts(p model.Product) (model.Product, error) {
	pI, err := pu.repository.CreateProduct(p)

	if err != nil {
		return model.Product{}, err
	}

	p.ID = pI

	return p, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil

}
