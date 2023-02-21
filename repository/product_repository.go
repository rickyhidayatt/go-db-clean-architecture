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

func (st *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := st.db.Select(&products, utils.SELECT_PRODUCT_LIST)
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

func NewProductRepository(dbStore *sqlx.DB) ProductRepository {
	return &productRepository{
		db: dbStore,
	}
}
