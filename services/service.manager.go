package services

import "github.com/jackc/pgx/v5/pgxpool"

type ServiceManager struct {
	*categoryService
}

func NewServiceManager(dbConn *pgxpool.Conn) *ServiceManager {
	return &ServiceManager{
		categoryService: NewCategoryService(dbConn),
	}
}
