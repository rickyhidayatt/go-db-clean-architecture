package model

type Product struct {
	Id      string `json:"id"`
	Name    string `json:"name" binding:"required"`
	Stock   int    `json:"stock" binding:"required"`
	Price   int    `json:"price" binding:"required"`
	StoreId string `db:"store_id"`
}
