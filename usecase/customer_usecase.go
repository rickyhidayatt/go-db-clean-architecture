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
	GetCustomerById(id string) (model.Customer, error)
	UpdateCustomerById(customer model.Customer) error
	DeleteCustomerById(id string) error
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerUseCase) RegisterCustomer(newCustomer *model.Customer) error {
	if len(newCustomer.Name) < 3 || len(newCustomer.Name) > 20 {
		return errors.New("Nama Minimal 3 sampai 20 Karakter")
	} else if newCustomer.Balance < 50000 {
		return errors.New("Minimal Saldo 50.000")
	} else {
		return c.customerRepo.Insert(newCustomer)
	}
}

func (c *customerUseCase) GetCustomerById(id string) (model.Customer, error) {
	return c.customerRepo.GetById(id)
}

func (c *customerUseCase) GetAllCustomer() ([]model.Customer, error) {
	return c.customerRepo.GetAll()
}

func (c *customerUseCase) UpdateCustomerById(customer model.Customer) error {
	if len(customer.Name) < 3 || len(customer.Name) > 20 {
		return errors.New("Nama Minimal 3 sampai 20 Karakter")
	} else if customer.Balance < 50000 {
		return errors.New("Minimal Saldo 50.000")
	} else {
		return c.customerRepo.UpdateById(customer)
	}
}

func (c *customerUseCase) DeleteCustomerById(id string) error {
	return c.customerRepo.DeleteById(id)
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepository,
	}
}
