package utils

const (
	INSERT_CUSTOMER       = "INSERT INTO customer (name, balance) VALUES ($1, $2)"
	SELECT_CUSTOMER_LIST  = "SELECT * FROM customer"
	SELECT_CUSTOMER_BY_ID = "SELECT * FROM customer WHERE id = $1"
	UPDATE_CUSTOMER       = "UPDATE customer SET name =$2, balance =$3 WHERE id = $1"
	DELETE_CUSTOMER       = "DELETE FROM customer WHERE id = $1"
)
