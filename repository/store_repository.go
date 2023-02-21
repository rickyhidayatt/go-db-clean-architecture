package repository

import (
	"database/sql"

	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type StoreRepository interface {
	GetAll() ([]model.Store, error)
	// GetAll() ([]model.StoreProducts, error)
}

type storeRepository struct {
	db *sql.DB
}

func (st *storeRepository) GetAll() ([]model.Store, error) {
	rows, err := st.db.Query(utils.SELECT_STORE_LIST)

	if err != nil {
		return nil, err
	}

	var stores []model.Store
	for rows.Next() {
		var store model.Store

		err := rows.Scan(&store.Id, &store.SiupNumber, &store.Name)

		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}
	return stores, nil
}

func NewStoreRepository(dbStore *sql.DB) StoreRepository {
	return &storeRepository{
		db: dbStore,
	}
}

// func (st *storeRepository) GetAll() ([]model.StoreProducts, error) {
// 	rows, err := st.db.Query(utils.SELECT_STORE_LIST)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var stores []model.StoreProducts

// 	for rows.Next() {
// 		var store model.StoreProducts

// 		// hardcode
// 		var siupNumber string

// 		err := rows.Scan(&store.Id, &siupNumber, &store.StoreName)

// 		if err != nil {
// 			return nil, err
// 		}

// 		var products []model.Product
// 		rows, errProd := st.db.Query(utils.SELECT_PRODUCT_STOREID, store.Id)
// 		if errProd != nil {
// 			return nil, err
// 		}

// 		for rows.Next() {
// 			var product model.Product
// 			err := rows.Scan(
// 				&product.Id,
// 				&product.Name,
// 				&product.Price,
// 				&product.Stock,
// 				&product.CreatedAt,
// 				&product.ProductId,
// 			)
// 			if err != nil {
// 				return nil, err
// 			}
// 			products = append(products, product)
// 		}
// 		store.Products = products
// 		stores = append(stores, store)
// 	}
// 	return stores, nil
// }
