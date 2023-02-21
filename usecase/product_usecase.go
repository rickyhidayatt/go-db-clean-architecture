package usecase

import (
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
)

type ProductUseCase interface {
	GetAllProduct() ([]model.Product, error)
	GetProductByStoreId(storeId string) ([]model.Product, error)
}

type productUseCase struct {
	productRepo repository.ProductRepository
}

func (pr *productUseCase) GetAllProduct() ([]model.Product, error) {
	return pr.productRepo.GetAll()
}

func (p *productUseCase) GetProductByStoreId(storeId string) ([]model.Product, error) {
	return p.productRepo.GetByStoreId(storeId)
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepo: productRepository,
	}
}
