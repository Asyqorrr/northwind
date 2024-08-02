-- name: FindCartsbyId :one
SELECT cart_id, customer_id, product_id, 
unit_price, qty, cart_created_on
FROM carts WHERE cart_id =$1;

-- name: FindAllCarts :many
SELECT cart_id, customer_id, product_id, 
unit_price, qty, cart_created_on
FROM carts;

-- name: FindAllCartsPaging :many
SELECT cart_id, customer_id, product_id, 
unit_price, qty, cart_created_on
FROM carts
ORDER BY cart_id
LIMIT $1 OFFSET $2;	

-- name: CreateCarts :one
INSERT INTO carts(customer_id, product_id, 
    unit_price, qty, cart_created_on)
VALUES ($1, $2, $3, $4, $5)
	RETURNING *;

-- name: UpdateCarts :one
UPDATE carts
	SET  customer_id=$2, product_id=$3, unit_price=$4, 
	qty=$5, cart_created_on=$6
WHERE cart_id=$1
RETURNING *;

-- name: DeleteCarts :exec
DELETE FROM carts
	WHERE cart_id=$1
    RETURNING *;

	