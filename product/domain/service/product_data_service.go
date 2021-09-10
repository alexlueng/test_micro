package service

import (
	"micro/product/domain/model"
	"micro/product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

func NewProductDataService(rp repository.IProductRepository) IProductDataService {
	return &ProductDataService{
		ProductRepository: rp,
	}
}

func (p *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return p.ProductRepository.CreateProduct(product)
}
func (p *ProductDataService) DeleteProduct(id int64) error {
	return p.ProductRepository.DeleteProductByID(id)
}
func (p *ProductDataService) UpdateProduct(product *model.Product) error {
	return p.ProductRepository.UpdateProduct(product)
}
func (p *ProductDataService) FindProductByID(id int64) (*model.Product, error) {
	return p.ProductRepository.FindProductByID(id)
}
func (p *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return p.ProductRepository.FindAll()
}
