package repository

import (
	"database/sql"

	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func (st *productRepository) GetAll() ([]model.Product, error) {
	rows, err := st.db.Query(utils.SELECT_PRODUCT_LIST)

	if err != nil {
		return nil, err
	}

	var products []model.Product
	for rows.Next() {
		var produk model.Product

		err := rows.Scan(&produk.Id, &produk.Name, &produk.Stock, &produk.Price, &produk.CreatedAt, &produk.ProductId)

		if err != nil {
			return nil, err
		}

		products = append(products, produk)
	}
	return products, nil
}

func NewProductRepository(dbStore *sql.DB) ProductRepository {
	return &productRepository{
		db: dbStore,
	}
}
