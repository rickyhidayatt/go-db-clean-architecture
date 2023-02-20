package usecase

import (
	"errors"

	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/repository"
)

// package ini buat validasi flow bisnis dari mana ke mana
// logic bisnis nya apa aja kaya nama di menu registrasi musti < 3 dll

type CustomerUseCase interface {
	RegisterCustomer(newCustomer *model.Customer) error
	GetAllCustomer() ([]model.Customer, error)
	GetCustomerByID(id int) (model.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerUseCase) RegisterCustomer(newCustomer *model.Customer) error {
	if len(newCustomer.Name) < 3 || len(newCustomer.Name) > 20 {
		return errors.New("Nama Minimal 3 sampai 20")
	} else if newCustomer.Balance < 50000 {
		return errors.New("Minimal 50.000")
	} else {
		return c.customerRepo.Insert(newCustomer)
	}
}

func (c *customerUseCase) GetAllCustomer() ([]model.Customer, error) {
	return c.customerRepo.GetAll()
}

func (c *customerUseCase) GetCustomerByID(id int) (model.Customer, error) {
	return c.customerRepo.GetCustomerByID(id)
}

func NewCustomerUseCase(customerRepositori repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepositori,
	}
}
