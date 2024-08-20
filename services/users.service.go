package services

import (
	db "courseonline/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersService struct {
	*db.Queries
}

// constructor
func NewUsersService(dbConn *pgxpool.Conn) *UsersService {
	return &UsersService{
		Queries: db.New(dbConn),
	}
}
