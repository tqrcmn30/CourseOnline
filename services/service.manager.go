package services

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceManager struct {
	*CartService
}

func NewServiceManager(dbConn *pgxpool.Conn) *ServiceManager {
	return &ServiceManager{
		CartService: NewCartService(dbConn),
	}

}
