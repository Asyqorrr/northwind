package services

import (
	db "b30northwindapi/db/sqlc"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService struct {
	*db.Queries // implements Querier
	dbConn      *pgxpool.Conn
}

func NewOrderService(dbConn *pgxpool.Conn) *OrderService {
	return &OrderService{
		Queries: db.New(dbConn),
		dbConn:  dbConn,
	}
}

func (order *OrderService) CreateOrderTx(ctx context.Context, args db.CreateOrderParams) (*db.Order, error) {
	tx, err := order.dbConn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	qtx := order.Queries.WithTx(tx)

	//populate cart list
	carts, err := qtx.FindCartByCustomerId(ctx, *args.CustomerID)
	if err != nil {
		return nil, err
	}

	log.Println(carts)

	//create order
	newOrder, err := qtx.CreateOrder(ctx, args)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(context.Background()); err != nil {
		return nil, err
	}
	return newOrder, nil

}
