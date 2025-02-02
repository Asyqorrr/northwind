// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: carts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCarts = `-- name: CreateCarts :one
INSERT INTO carts(customer_id, product_id, 
    unit_price, qty, cart_created_on)
VALUES ($1, $2, $3, $4, $5)
	RETURNING cart_id, customer_id, product_id, unit_price, qty, cart_created_on
`

type CreateCartsParams struct {
	CustomerID    string      `json:"customer_id"`
	ProductID     int32       `json:"product_id"`
	UnitPrice     *float32    `json:"unit_price"`
	Qty           *int32      `json:"qty"`
	CartCreatedOn pgtype.Date `json:"cart_created_on"`
}

func (q *Queries) CreateCarts(ctx context.Context, arg CreateCartsParams) (*Cart, error) {
	row := q.db.QueryRow(ctx, createCarts,
		arg.CustomerID,
		arg.ProductID,
		arg.UnitPrice,
		arg.Qty,
		arg.CartCreatedOn,
	)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.CustomerID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Qty,
		&i.CartCreatedOn,
	)
	return &i, err
}

const deleteCarts = `-- name: DeleteCarts :exec
DELETE FROM carts
	WHERE cart_id=$1
    RETURNING cart_id, customer_id, product_id, unit_price, qty, cart_created_on
`

func (q *Queries) DeleteCarts(ctx context.Context, cartID int32) error {
	_, err := q.db.Exec(ctx, deleteCarts, cartID)
	return err
}

const findAllCarts = `-- name: FindAllCarts :many
SELECT cart_id, customer_id, product_id, 
unit_price, qty, cart_created_on
FROM carts
`

func (q *Queries) FindAllCarts(ctx context.Context) ([]*Cart, error) {
	rows, err := q.db.Query(ctx, findAllCarts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Cart
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.CartID,
			&i.CustomerID,
			&i.ProductID,
			&i.UnitPrice,
			&i.Qty,
			&i.CartCreatedOn,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllCartsPaging = `-- name: FindAllCartsPaging :many
SELECT cart_id, customer_id, product_id, 
unit_price, qty, cart_created_on
FROM carts
ORDER BY cart_id
LIMIT $1 OFFSET $2
`

type FindAllCartsPagingParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) FindAllCartsPaging(ctx context.Context, arg FindAllCartsPagingParams) ([]*Cart, error) {
	rows, err := q.db.Query(ctx, findAllCartsPaging, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Cart
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.CartID,
			&i.CustomerID,
			&i.ProductID,
			&i.UnitPrice,
			&i.Qty,
			&i.CartCreatedOn,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findCartsbyId = `-- name: FindCartsbyId :one
SELECT cart_id, customer_id, product_id, 
unit_price, qty, cart_created_on
FROM carts WHERE cart_id =$1
`

func (q *Queries) FindCartsbyId(ctx context.Context, cartID int32) (*Cart, error) {
	row := q.db.QueryRow(ctx, findCartsbyId, cartID)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.CustomerID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Qty,
		&i.CartCreatedOn,
	)
	return &i, err
}

const updateCarts = `-- name: UpdateCarts :one
UPDATE carts
	SET  customer_id=$2, product_id=$3, unit_price=$4, 
	qty=$5, cart_created_on=$6
WHERE cart_id=$1
RETURNING cart_id, customer_id, product_id, unit_price, qty, cart_created_on
`

type UpdateCartsParams struct {
	CartID        int32       `json:"cart_id"`
	CustomerID    string      `json:"customer_id"`
	ProductID     int32       `json:"product_id"`
	UnitPrice     *float32    `json:"unit_price"`
	Qty           *int32      `json:"qty"`
	CartCreatedOn pgtype.Date `json:"cart_created_on"`
}

func (q *Queries) UpdateCarts(ctx context.Context, arg UpdateCartsParams) (*Cart, error) {
	row := q.db.QueryRow(ctx, updateCarts,
		arg.CartID,
		arg.CustomerID,
		arg.ProductID,
		arg.UnitPrice,
		arg.Qty,
		arg.CartCreatedOn,
	)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.CustomerID,
		&i.ProductID,
		&i.UnitPrice,
		&i.Qty,
		&i.CartCreatedOn,
	)
	return &i, err
}
