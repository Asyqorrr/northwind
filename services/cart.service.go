package services

import (
	db "b30northwindapi/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CartService struct {
	*db.Queries
}

// constructor
func NewCartService(dbConn *pgxpool.Conn) *CartService {
	return &CartService{
		Queries: db.New(dbConn),
	}
}
