package model

type Store struct {
	Id         string
	SiupNumber string `db:"siup_number"`
	Name       string
}
