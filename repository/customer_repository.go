package repository

// package repository bertanggung jawab untuk CRUD Ke Database

import (
	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type CustomerRepository interface {
	Insert(newCustomer *model.Customer) error
	GetAll() ([]model.Customer, error)
	GetCustomerByID(id string) (model.Customer, error)
	Update(updateId string, updateCustomer *model.Customer) error
	Delete(id string) error
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Insert(newCustomer *model.Customer) error {
	// func ini buat memenuhi kontrak
	// defer c.db.Close()
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, newCustomer)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(updateId string, updateCustomer *model.Customer) error {
	_, err := c.db.Exec(utils.UPDATE_CUSTOMER, updateId, updateCustomer.Name, updateCustomer.Balance)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Delete(id string) error {
	_, err := c.db.Exec(utils.DELETE_CUSTOMER, id)
	if err != nil {
		return err
	}

	return nil

}

func (c *customerRepository) GetAll() ([]model.Customer, error) {

	var customers []model.Customer // menampung dari looping di bawah
	err := c.db.Select(&customers, utils.SELECT_CUSTOMER_LIST)

	if err != nil {
		return nil, err
	}
	return customers, nil

}

func (c *customerRepository) GetCustomerByID(id string) (model.Customer, error) {
	var customer model.Customer

	rows := c.db.QueryRow(utils.SELECT_CUSTOMER_BY_ID, id)

	err := rows.Scan(&customer.Id, &customer.Name, &customer.Balance)
	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil

}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
