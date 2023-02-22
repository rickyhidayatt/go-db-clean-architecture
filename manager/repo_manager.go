package manager

import "github.com/rickyhidayatt/repository"

type RepositoryManager interface {
	ProductRepository() repository.ProductRepository
	StoreRepository() repository.StoreRepository
	CustomerRepository() repository.CustomerRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) ProductRepository() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}

func (r *repositoryManager) StoreRepository() repository.StoreRepository {
	return repository.NewStoreRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infraManager,
	}
}
