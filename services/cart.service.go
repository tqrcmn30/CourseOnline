package services

import (
	db "courseonline/db/sqlc"

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
