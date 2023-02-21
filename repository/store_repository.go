package repository

import (
	// "database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/model"
	"github.com/rickyhidayatt/utils"
)

type StoreRepository interface {
	GetAll() ([]model.Store, error)
}

type storeRepository struct {
	db *sqlx.DB
}

func (st *storeRepository) GetAll() ([]model.Store, error) {
	// Tempat untuk menampung slice dengan tipe data model store
	var stores []model.Store
	err := st.db.Select(&stores, utils.SELECT_STORE_LIST)
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func NewStoreRepository(dbStore *sqlx.DB) StoreRepository {
	return &storeRepository{
		db: dbStore,
	}
}
