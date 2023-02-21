package repository

import (
	"database/sql"

	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type JoinRepository interface {
	GetAllProduckAndStore() ([]model.Store, []model.Product, error)
	// GetAll() ([]model.StoreProducts, error)
}

type joinRepository struct {
	db *sql.DB
}

func (st *joinRepository) GetAllProduckAndStore() ([]model.Store, []model.Product, error) {
	rows, err := st.db.Query(utils.JOIN_PRODUCT_STORE)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var stores []model.Store
	var products []model.Product

	for rows.Next() {
		var store model.Store
		var product model.Product

		err := rows.Scan(&store.Id, &store.SiupNumber, &store.Name, &product.Id, &product.Name, &product.Stock, &product.Price, &product.CreatedAt, &product.ProductId)

		if err != nil {
			return nil, nil, err
		}

		stores = append(stores, store)
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return stores, products, nil
}

func JoinStoreProdukRepository(dbStore *sql.DB) JoinRepository {
	return &joinRepository{
		db: dbStore,
	}
}
