package model

type Product struct {
	Id        string
	Name      string
	Stock     int
	Price     int
	CreatedAt string `db:"created_at"`
	StoreId   string `db:"store_id"`
}
