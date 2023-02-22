package usecase

import (
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
	"github.com/rickyhidayatt/utils"
)

type ProductUseCase interface {
	CreateNewProduct(newProduct *model.Product) error
	GetAllProduct() ([]model.Product, error)
	GetProductByStoreId(storeId string) ([]model.Product, error)
}

type productUseCase struct {
	productRepo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(newProduct *model.Product) error {
	newProduct.Id = utils.GenerateId()
	return p.productRepo.Insert(*newProduct)
}

func (p *productUseCase) GetAllProduct() ([]model.Product, error) {

	return p.productRepo.GetAll()
}

func (p *productUseCase) GetProductByStoreId(storeId string) ([]model.Product, error) {
	return p.productRepo.GetByStoreId(storeId)
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepo: productRepository,
	}
}
