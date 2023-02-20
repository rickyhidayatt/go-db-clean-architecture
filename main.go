package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Repsol12"
	dbname   = "incubation_class"
)

type Customer struct {
	Id      string
	Name    string
	Balance int
}

var connectDB = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	db, err := sql.Open("postgres", connectDB)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	newCustomer := Customer{
		Name:    "Dicky Tanda tanya",
		Balance: 1029292,
	}

	insertData(db, newCustomer)
}

func insertData(db *sql.DB, newCustomer Customer) {
	query := "INSERT INTO customer (name, balance) VALUES ($1, $2)"
	result, err := db.Exec(query, newCustomer.Name, newCustomer.Balance)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
