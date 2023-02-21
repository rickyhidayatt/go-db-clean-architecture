package usecase

import (
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
	"github.com/rickyhidayatt/utils"
)

type StoreUseCase interface {
	GetAllStore() ([]model.Store, error)
	GetAllStoreWithProducts() ([]model.StoreProducts, error)
}

type storeUseCase struct {
	storeRepo repository.StoreRepository
	productUc ProductUseCase
}

func (st *storeUseCase) GetAllStore() ([]model.Store, error) {
	return st.storeRepo.GetAll()
}

func (st *storeUseCase) GetAllStoreWithProducts() ([]model.StoreProducts, error) {
	var storeProducts []model.StoreProducts
	stores, err := st.GetAllStore()
	if utils.ErrorCheck(err) {
		return nil, err
	}
	for _, store := range stores {
		var storeProduct model.StoreProducts
		storeProduct.Id = store.Id
		storeProduct.StoreName = store.Name
		products, err := st.productUc.GetProductByStoreId(store.Id)
		if err != nil {
			return nil, err
		}
		storeProduct.Products = products
		storeProducts = append(storeProducts, storeProduct)
	}
	return storeProducts, nil
}

func NewStoreUseCase(storeRepository repository.StoreRepository, productUseCase ProductUseCase) StoreUseCase {
	return &storeUseCase{
		storeRepo: storeRepository,
		productUc: productUseCase,
	}
}
