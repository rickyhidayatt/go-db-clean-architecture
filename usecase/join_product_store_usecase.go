package usecase

import (
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
)

type JoinUseCase interface {
	JoinProductAndStore() ([]model.Store, []model.Product, error)
}

type joinUseCase struct {
	joinproduct repository.JoinRepository
}

func (pr *joinUseCase) JoinProductAndStore() ([]model.Store, []model.Product, error) {
	return pr.joinproduct.GetAllProduckAndStore()
}

func JoinProductUseCase(joinRepo repository.JoinRepository) JoinUseCase {
	return &joinUseCase{
		joinproduct: joinRepo,
	}
}
