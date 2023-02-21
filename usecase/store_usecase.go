package usecase

import (
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
)

type StoreUseCase interface {
	GetAllStore() ([]model.Store, error)
	// GetAllStore() ([]model.StoreProducts, error)
}

type storeUseCase struct {
	storeRepo repository.StoreRepository
}

func (st *storeUseCase) GetAllStore() ([]model.Store, error) {
	return st.storeRepo.GetAll()
}

// func (st *storeUseCase) GetAllStore() ([]model.StoreProducts, error) {
// 	return st.storeRepo.GetAll()
// }

func NewStoreUseCase(storeRepository repository.StoreRepository) StoreUseCase {
	return &storeUseCase{
		storeRepo: storeRepository,
	}
}
