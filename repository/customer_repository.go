package repository

// package repository bertanggung jawab untuk

import (
	"database/sql"

	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type CustomerRepository interface {
	Insert(newCustomer *model.Customer) error
	GetAll() ([]model.Customer, error)
	GetCustomerByID(id int) (model.Customer, error)
	Update(updateId string, updateCustomer *model.Customer) error
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Insert(newCustomer *model.Customer) error {
	// func ini buat memenuhi kontrak
	// defer c.db.Close()
	_, err := c.db.Exec(utils.INSERT_CUSTOMER, newCustomer.Name, newCustomer.Balance)
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

func (c *customerRepository) GetAll() ([]model.Customer, error) {
	rows, err := c.db.Query(utils.SELECT_CUSTOMER_LIST)
	if err != nil {
		return nil, err
	}
	var customers []model.Customer // menampung dari looping di bawah
	for rows.Next() {
		var customer model.Customer // apa yang mau di looping dari db

		err := rows.Scan(&customer.Id, &customer.Name, &customer.Balance)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil

}

func (c *customerRepository) GetCustomerByID(id int) (model.Customer, error) {
	var customer model.Customer

	rows := c.db.QueryRow(utils.SELECT_CUSTOMER_BY_ID, id)

	err := rows.Scan(&customer.Id, &customer.Name, &customer.Balance)
	if err != nil {
		return customer, err
	}

	return customer, nil

}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
