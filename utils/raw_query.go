package utils

const (
	INSERT_CUSTOMER      = "INSERT INTO customer (name, balance) VALUES ($1, $2)"
	SELECT_CUSTOMER_LIST = "SELECT * FROM customers"

	SELECT_CUSTOMER_BY_ID  = "SELECT * FROM customer WHERE id = $1"
	UPDATE_CUSTOMER        = "UPDATE customer SET name =$2, balance =$3 WHERE id = $1"
	DELETE_CUSTOMER        = "DELETE FROM customer WHERE id = $1"
	SELECT_STORE_LIST      = "SELECT * FROM mst_store"
	SELECT_PRODUCT_LIST    = "SELECT * FROM mst_product"
	SELECT_PRODUCT_STOREID = "SELECT * FROM mst_product WHERE store_id = $1"
	// join
	JOIN_PRODUCT_STORE = `
	SELECT s.id, s.siup_number, s.name, p.id AS product_id, p.name, p.stock, p.price, p.created_at, p.store_id
FROM mst_store s
JOIN mst_product p ON s.id = p.store_id;
`
)
