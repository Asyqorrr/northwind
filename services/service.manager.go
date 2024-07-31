package services

import "github.com/jackc/pgx/v5/pgxpool"

type ServiceManager struct {
	*CategoryService
	*ProductService
}

func NewServiceManager(dbConn *pgxpool.Conn) *ServiceManager {
	return &ServiceManager{
		CategoryService: NewCategoryService(dbConn),
		ProductService: NewProductService(dbConn),
	}
}
