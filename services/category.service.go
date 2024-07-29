package services

import (
	db "b30northwindapi/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryService struct {
	*db.Queries
}

// constructor
func NewCategoryService(dbConn *pgxpool.Conn) *categoryService {
	return &categoryService{
		Queries: db.New(dbConn),
	}
}
