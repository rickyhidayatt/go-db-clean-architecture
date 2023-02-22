package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	GetByStoreId(id string) ([]model.Product, error)
}

type productRepository struct {
	db *sqlx.DB
}

func (p *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Select(&products, utils.SELECT_PRODUCT_LIST)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepository) GetByStoreId(id string) ([]model.Product, error) {
	var products []model.Product
	err := p.db.Select(&products, utils.SELECT_PRODUCT_STOREID, id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
