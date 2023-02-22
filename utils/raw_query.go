package utils

const (
	//PRODUCT QUERY
	INSERT_PRODUCT      = "INSERT INTO product (id,name,price,stock) VALUES (:id,:name,:price,:stock)"
	SELECT_PRODUCT_LIST = "SELECT id, name, price, stock FROM product"
	UPDATE_PRODUCT      = "UPDATE product SET name =$2, price =$3, stock =$4, store_id=$5 WHERE id = $1"

	//CUSTOMER QUERY
	INSERT_CUSTOMER      = "INSERT INTO customer (name, balance) VALUES ($1, $2)"
	SELECT_CUSTOMER_LIST = "SELECT * FROM customers"
	SELECT_CUSTOMER_ID   = "SELECT * FROM customer WHERE id = $1"
	UPDATE_CUSTOMER      = "UPDATE customer SET name =$2, balance =$3 WHERE id = $1"
	DELETE_CUSTOMER      = "DELETE FROM customer WHERE id = $1"

	//STORE QUERY
	SELECT_STORE_LIST      = "SELECT * FROM mst_store"
	SELECT_PRODUCT_STOREID = "SELECT * FROM mst_product WHERE store_id = $1"
)
