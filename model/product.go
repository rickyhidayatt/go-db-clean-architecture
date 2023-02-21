package model

import "time"

type Product struct {
	Id        string
	Name      string
	Stock     int
	Price     int
	CreatedAt time.Time
	ProductId string
}
