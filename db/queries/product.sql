-- name: FindProductById :one
SELECT product_id, product_name, supplier_id, 
category_id, quantity_per_unit, unit_price, 
units_in_stock, units_on_order, reorder_level, discontinued,
product_image 
FROM products WHERE product_id =$1;

-- name: FindAllProduct :many
SELECT product_id, product_name, supplier_id, 
category_id, quantity_per_unit, unit_price, 
units_in_stock, units_on_order, reorder_level, discontinued,product_image
	FROM products;

-- name: FindAllProductPaging :many
SELECT product_id, product_name, supplier_id, 
category_id, quantity_per_unit, unit_price, 
units_in_stock, units_on_order, reorder_level, discontinued,product_image
FROM products
ORDER BY product_id
LIMIT $1 OFFSET $2;	

-- name: CreateProduct :one
INSERT INTO products(
	product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued,product_image)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10)
	RETURNING *;

-- name: UpdateProduct :one
UPDATE products
	SET  product_name=$2, supplier_id=$3, category_id=$4, 
	quantity_per_unit=$5, unit_price=$6, units_in_stock=$7, 
	units_on_order=$8, reorder_level=$9, discontinued=$10,product_image=$11
	WHERE product_id=$1
	RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
	WHERE product_id=$1
    RETURNING *;

	