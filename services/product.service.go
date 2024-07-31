package services

import (
	db "b30northwindapi/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	*db.Queries
}

// constructor
func NewProductService(dbConn *pgxpool.Conn) *ProductService {
	return &ProductService{
		Queries: db.New(dbConn),
	}
}
