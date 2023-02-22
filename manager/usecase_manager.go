package manager

import "github.com/rickyhidayatt/usecase"

type UseCaseManager interface {
	ProductUseCase() usecase.ProductUseCase
	StoreUseCase() usecase.StoreUseCase
	CustomerUseCase() usecase.CustomerUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepository())
}

func (u *useCaseManager) StoreUseCase() usecase.StoreUseCase {
	return usecase.NewStoreUseCase(
		u.repoManager.StoreRepository(),
		u.ProductUseCase(),
	)
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepository())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
