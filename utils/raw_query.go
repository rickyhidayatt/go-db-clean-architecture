package utils

const (
	INSERT_CUSTOMER       = "INSERT INTO customer (name, balance) VALUES ($1, $2)"
	SELECT_CUSTOMER_LIST  = "SELECT * FROM customer"
	SELECT_CUSTOMER_BY_ID = "SELECT * FROM customer WHERE id = $1"
)
